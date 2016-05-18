package anagrambler

import (
	"bytes"
	"io/ioutil"
)

func sortedLower(word []byte) []byte {
	sorted := bytes.ToLower(word)

	for i := 0; i < len(word); i++ {
		c := sorted[i]
		j := i - 1
		for ; j >= 0 && sorted[j] > c; j-- {
			sorted[j+1] = sorted[j]
		}
		sorted[j+1] = c
	}

	return sorted
}

func (trie *Trie) LoadDict(filepath string) error {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	words := bytes.Split(data, []byte("\n"))
	words = words[:len(words)-1]

	for _, word := range words {
		trie.AddWord(word)
	}

	return nil
}

func (trie *Trie) AddWord(word []byte) {
	path := trie.Root

	for _, letter := range sortedLower(word) {
		if path.Children[letter - 'a'] == nil {
			path.Children[letter - 'a'] = trie.NextNode()
		}
		path = path.Children[letter - 'a']
	}

	// Add to the head of the linked list of anagrams
	w := &Word{s: word, next: path.Words}

	path.Words = w
}

func (trie *Trie) Search(text string, filter string) []string {
	results := make(map[*Node]bool)

	t, f := []byte(text), []byte(filter)

	search(trie.Root, sortedLower(t), sortedLower(f), results)

	filteredResults := make([]string, 0)

	for node := range results {
		for word := node.Words; word != nil; word = word.next {
			if bytes.Contains(word.s, f) {
				filteredResults = append(filteredResults, string(word.s))
			}
		}
	}

	return filteredResults
}

func search(n *Node, text []byte, filter []byte, results map[*Node]bool) {
	// Record any words stored at this node
	// Only record acronyms after the filter has been satisfied
	if len(filter) == 0 && n.Words != nil {
		if !results[n] {
			// Add this node's acronyms to the results
			results[n] = true
		} else {
			// We've already traversed this node, so stop searching it
			return
		}
	}

	// Keep track of which runes we've searched
	searched_runes := make(map[byte]bool)

	for i, letter := range text {
		// Skip any runes that we don't have nodes for
		// or that we've already searched for (i.e. duplicate runes)
		if n.Children[letter - 'a'] == nil || searched_runes[letter] == true {
			continue
		}

		var new_filter []byte

		switch {
		case len(filter) == 0:
			// The filter has already been satisfied
			new_filter = filter
		case letter < filter[0]:
			// This letter doesn't affect the filter
			new_filter = filter
		case letter == filter[0]:
			// This letter satisfies the next rune in the filter, so we can
			// remove it from the filter
			new_filter = filter[1:]
		case letter > filter[0]:
			// The remaining letters in the text are all greater than the next
			// required filter rune, so none of the remaining substrings will
			// satisfy the filter
			return
		}

		search(n.Children[letter - 'a'], text[i+1:], new_filter, results)

		searched_runes[letter] = true
	}
}
