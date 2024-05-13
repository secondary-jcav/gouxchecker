package main

import (
	"fmt"
	"time"

	"github.com/secondary-jcav/gouxchecker/pkg/images"
	"github.com/secondary-jcav/gouxchecker/pkg/output"
	"github.com/secondary-jcav/gouxchecker/pkg/scraper"
	"github.com/secondary-jcav/gouxchecker/pkg/spelling"
)

func main() {
	// create an url without trailing slashes
	url := "http://localhost:8080"
	start := time.Now()
	c := scraper.InitializeCollector(url)
	s := spelling.InitSpellchecker()

	fmt.Println("Scraping target website")

	// Start the scraping process and receive the collected data
	fontSet, altTexts, misspelledWords := scraper.StartScraping(c, s, url)

	noAlts, duplicateAlts := images.CheckAltText(altTexts)

	output.ImageAlts(noAlts, duplicateAlts)
	output.FontsResults(fontSet)
	output.Typos(misspelledWords)

	elapsed := time.Since(start)
	fmt.Println("Finished after ", elapsed)
}
