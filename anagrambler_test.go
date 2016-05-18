package anagrambler_test

import (
	"testing"

	"github.com/RyanEdwardHall/anagrambler"
)

func TestKnownOutput(t *testing.T) {
	trie := anagrambler.NewTrie()

	if err := trie.LoadDict("go-dict.txt"); err != nil {
		t.Error("Could not load dictionary 'go-dict.txt'", err)
	}

	searchWord := "honorificabilitudinitatibus"

	results := trie.Search(searchWord, "")

	if len(results) != 9083 {
		t.Error("Expected 9083 words, got ", len(results))
	}
}

func BenchmarkAnagrambler(b *testing.B) {
	trie := anagrambler.NewTrie()

	if err := trie.LoadDict("go-dict.txt"); err != nil {
		b.Error("Could not load dictionary 'go-dict.txt'", err)
	}

	searchWord := "Lopadotemachoselachogaleokranioleipsanodrimhypotrimmatosilphioparaomelitokatakechymenokichlepikossyphophattoperisteralektryonoptekephalliokigklopeleiolagoiosiraiobaphetraganopterygon"

	b.ResetTimer()

	for counter := 0; counter < b.N; counter++ {
		trie.Search(searchWord, "")
	}
}
