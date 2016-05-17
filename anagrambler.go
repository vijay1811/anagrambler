package anagrambler

import (
	"io/ioutil"
	"sort"
	"strings"
)

func sortedLower(w string) string {
	w = strings.ToLower(w)
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func LoadDict(trie *Trie, filepath string) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	words := strings.Split(string(data), "\n")
	words = words[:len(words)-1]

	for _, word := range words {
		AddWord(trie, word)
	}
}

func AddWord(trie *Trie, word string) {
	path := trie.Root

	for _, letter := range sortedLower(word) {
		if path.getChild(letter) == nil {
			path.setChild(letter, NextNode(trie))
		}
		path = path.getChild(letter)
	}

	// Add to the head of the linked list of anagrams
	w := &Word{s: &word, next: path.Words}

	path.Words = w
}

func Search(trie *Trie, text string, filter string) []string {
	results := make(map[*Node]bool)

	search(trie.Root, sortedLower(text), sortedLower(filter), results)

	filteredResults := make([]string, 0)

	for node := range results {
		for word := node.Words; word != nil; word = word.next {
			if strings.Contains(*word.s, filter) {
				filteredResults = append(filteredResults, *word.s)
			}
		}
	}

	return filteredResults
}

func search(n *Node, text string, filter string, results map[*Node]bool) {
	// Record any words stored at this node
	// Only record acronyms after the filter has been satisfied
	if filter == "" && n.Words != nil {
		if !results[n] {
			// Add this node's acronyms to the results
			results[n] = true
		} else {
			// We've already traversed this node, so stop searching it
			return
		}
	}

	// Keep track of which runes we've searched
	searched_runes := make(map[rune]bool)

	for i, letter := range text {
		// Skip any runes that we don't have nodes for
		// or that we've already searched for (i.e. duplicate runes)
		if n.getChild(letter) == nil || searched_runes[letter] == true {
			continue
		}

		var new_filter string

		switch {
		case filter == "":
			// The filter has already been satisfied
			new_filter = ""
		case letter < rune(filter[0]):
			// This letter doesn't affect the filter
			new_filter = filter[:]
		case letter == rune(filter[0]):
			// This letter satisfies the next rune in the filter, so we can
			// remove it from the filter
			new_filter = filter[1:]
		case letter > rune(filter[0]):
			// The remaining letters in the text are all greater than the next
			// required filter rune, so none of the remaining substrings will
			// satisfy the filter
			return
		}

		search(n.getChild(letter), text[i+1:], new_filter, results)

		searched_runes[letter] = true
	}
}
