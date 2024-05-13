package scraper

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/secondary-jcav/gouxchecker/pkg/fonts"
	"github.com/secondary-jcav/gouxchecker/pkg/spelling"

	"github.com/gocolly/colly/v2"
	"github.com/makifdb/spellcheck"
)

func InitializeCollector(domain string) *colly.Collector {
	// Extract just the domain name, removing protocol and port
	u, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector(
		colly.AllowedDomains(u.Hostname()),
	)

	return c
}

func SetupHandlers(c *colly.Collector, fontSet map[string]bool, altTexts map[string][]string, wg *sync.WaitGroup, misspelledWords map[string]bool, sc *spellcheck.Trie, brokenLinks map[string]bool) {
	c.OnHTML(`link[rel="stylesheet"]`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		wg.Add(1) // Add to the WaitGroup just before initiating the goroutine
		go func() {
			defer wg.Done() // Ensure Done is called to signal completion
			e.Request.Visit(link)
		}()
	})

	c.OnHTML("style", func(e *colly.HTMLElement) {
		fonts.ExtractFonts(e.Text, fontSet)
	})

	c.OnResponse(func(r *colly.Response) {

		if strings.Contains(r.Headers.Get("Content-Type"), "text/css") {
			fonts.ExtractFonts(string(r.Body), fontSet)
		}

	})

	c.OnError(func(r *colly.Response, err error) {

		if r.StatusCode != http.StatusOK {
			brokenLinks[r.Request.URL.String()] = true
		}
	})

	c.OnHTML("img", func(e *colly.HTMLElement) {
		alt := e.Attr("alt")
		src := e.Request.AbsoluteURL(e.Attr("src"))
		altTexts[alt] = append(altTexts[alt], src)
	})

	c.OnHTML("p, h1, h2, h3, h4, h5, h6, li, span, a", func(e *colly.HTMLElement) {
		text := e.Text
		spelling.ExtractAndCheckText(text, sc, misspelledWords)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// makes sure we visit every new link inside the page
		link := e.Attr("href")
		absoluteURL := e.Request.AbsoluteURL(link)
		parsedURL, err := url.Parse(absoluteURL)
		if err != nil {
			log.Printf("Error parsing URL '%s': %v", absoluteURL, err)
			return
		}

		if parsedURL.String() == e.Request.URL.String() {
			// don't visit old links
			return
		}

		if parsedURL.Host == e.Request.URL.Host {
			c.Visit(parsedURL.String())

		}
	})

}

// StartScraping begins the scraping process on the specified URL
func StartScraping(c *colly.Collector, sc *spellcheck.Trie, url string) (map[string]bool, map[string][]string, map[string]bool, map[string]bool) {
	fontSet := make(map[string]bool)
	altTexts := make(map[string][]string)
	misspelledWords := make(map[string]bool)
	brokenLinks := make(map[string]bool)
	wg := &sync.WaitGroup{}
	SetupHandlers(c, fontSet, altTexts, wg, misspelledWords, sc, brokenLinks)
	c.Visit(url + "/")
	c.Wait()  // Wait for all collectors to complete, including async visits
	wg.Wait() // Wait for all goroutines to finish
	return fontSet, altTexts, misspelledWords, brokenLinks
}
