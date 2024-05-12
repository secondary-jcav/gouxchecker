package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/secondary-jcav/gouxchecker/pkg/images"
	"github.com/secondary-jcav/gouxchecker/pkg/scraper"
	"github.com/secondary-jcav/gouxchecker/pkg/spelling"
)

func main() {
	url := "http://localhost:8080/"
	start := time.Now()
	c := scraper.InitializeCollector(url)
	s := spelling.InitSpellchecker()

	// Start the scraping process and receive the collected data
	fontSet, altTexts, misspelledWords := scraper.StartScraping(c, s, url)

	fmt.Println("Fonts found:")
	for font := range fontSet {
		fmt.Println(font)
	}

	noAlts, duplicateAlts := images.CheckAltText(altTexts)
	fmt.Println("no alts:")
	fmt.Println(noAlts)

	fmt.Println("duplicateAlts")
	fmt.Println(duplicateAlts)

	if len(misspelledWords) != 0 {
		var keys []string
		for key := range misspelledWords {
			keys = append(keys, key)
		}
		fmt.Println("Possible typos found")
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key)
		}
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
