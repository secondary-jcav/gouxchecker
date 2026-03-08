package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	urlPtr := flag.String("url", "http://localhost:8080", "Target URL to check")
	flag.Parse()
	if *urlPtr == "" {
		log.Fatal("Please provide a URL using -url flag")
	}
	url := strings.TrimRight(*urlPtr, "/")

	if err := os.MkdirAll("results", 0755); err != nil {
		log.Fatal("Could not create results directory: ", err)
	}

	start := time.Now()
	c := scraper.InitializeCollector(url)
	s := spelling.InitSpellchecker()

	fmt.Println("Scraping target website")

	result := scraper.StartScraping(c, s, url)

	noAlts, duplicateAlts := images.CheckAltText(result.AltTexts)

	nonResponsivePages := responsive.CheckResponsiveUrls(result.VisitedUrls, result.HasViewport)

	output.ImageAlts(noAlts, duplicateAlts)
	output.FontsResults(result.FontSet)
	output.Typos(result.MisspelledWords)
	output.BrokenLinks(result.BrokenLinks)
	output.NonResponsivePages(result.ContainsMediaQueries, nonResponsivePages)

	elapsed := time.Since(start)
	fmt.Println("Finished after ", elapsed)
}
