package spelling

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/makifdb/spellcheck"
)

func InitSpellchecker() *spellcheck.Trie {
	sc, err := spellcheck.New()
	if err != nil {
		fmt.Println("spellchecker error init")
	}

	return sc
}

func ExtractAndCheckText(text string, sc *spellcheck.Trie, misspelledWords map[string]bool) {
	// Extract text and split into words
	reg := regexp.MustCompile(`[^\w\s]`)
	numericReg := regexp.MustCompile(`^\d+$`)
	words := strings.Fields(text)
	for _, word := range words {
		// Remove punctuation from the word and convert to lowercase
		cleanWord := reg.ReplaceAllString(word, "")
		lowerWord := strings.ToLower(cleanWord)
		// Check if the word is purely numeric or non-empty
		if lowerWord != "" && !numericReg.MatchString(lowerWord) {
			ok := sc.SearchDirect(lowerWord)
			if !ok {
				misspelledWords[lowerWord] = true
			}
		}
	}
}
