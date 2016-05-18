package main

import (
	"fmt"
	"os"

	"github.com/RyanEdwardHall/anagrambler"
)

func main() {
	trie := anagrambler.NewTrie()

	if err := trie.LoadDict("go-dict.txt"); err != nil {
		fmt.Printf("Could not load dictionary 'go-dict.txt'")
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		searchWord := os.Args[1]

		filter := ""

		if len(os.Args) == 3 {
			filter = os.Args[2]
		}

		results := trie.Search(searchWord, filter)

		fmt.Println("Number of anagrams:", len(results))

		for _, anagram := range results {
			fmt.Println(anagram)
		}
	} else {
		fmt.Println("No search string specified")
	}
}
