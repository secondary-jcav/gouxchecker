package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/secondary-jcav/gouxchecker/pkg/images"
	"github.com/secondary-jcav/gouxchecker/pkg/output"
	"github.com/secondary-jcav/gouxchecker/pkg/responsive"
	"github.com/secondary-jcav/gouxchecker/pkg/scraper"
	"github.com/secondary-jcav/gouxchecker/pkg/spelling"
)

func main() {
	// create a url without trailing slashes
	urlPtr := flag.String("url", "https://mariaalejandraperez.com/", "Target URL to check")
	flag.Parse()
	if *urlPtr == "" {
		log.Fatal("Please provide a URL using -url flag")
	}
	url := strings.TrimRight(*urlPtr, "/")
	start := time.Now()
	c := scraper.InitializeCollector(url)
	s := spelling.InitSpellchecker()

	fmt.Println("Scraping target website")

	// Start the scraping process and receive the collected data
	fontSet, altTexts, misspelledWords, brokenLinks, nonNativeUrls, visitedUrls, containsMediaQueries := scraper.StartScraping(c, s, url)

	noAlts, duplicateAlts := images.CheckAltText(altTexts)

	nonResponsivePages := responsive.CheckResponsiveUrls(visitedUrls, nonNativeUrls)

	output.ImageAlts(noAlts, duplicateAlts)
	output.FontsResults(fontSet)
	output.Typos(misspelledWords)
	output.BrokenLinks(brokenLinks)
	output.NonResponsivePages(containsMediaQueries, nonResponsivePages)

	elapsed := time.Since(start)
	fmt.Println("Finished after ", elapsed)
}
