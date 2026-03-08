package scraper

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/secondary-jcav/gouxchecker/pkg/fonts"
	"github.com/secondary-jcav/gouxchecker/pkg/responsive"
	"github.com/secondary-jcav/gouxchecker/pkg/spelling"

	"github.com/gocolly/colly/v2"
	"github.com/makifdb/spellcheck"
)

// ScrapeResult holds all data collected during a scraping run.
type ScrapeResult struct {
	FontSet              map[string]bool
	AltTexts             map[string][]string
	MisspelledWords      map[string]bool
	BrokenLinks          map[string]bool
	HasViewport          map[string]bool // true when a page has <meta name="viewport">
	VisitedUrls          map[string]bool
	ContainsMediaQueries bool
}

func InitializeCollector(domain string) *colly.Collector {
	// Extract just the domain name, removing protocol and port
	u, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector(
		colly.AllowedDomains(u.Hostname()),
	)
	c.SetRequestTimeout(30 * time.Second)

	return c
}

func SetupHandlers(c *colly.Collector, sc *spellcheck.Trie, wg *sync.WaitGroup, result *ScrapeResult, mu *sync.Mutex) {
	c.OnHTML(`link[rel="stylesheet"]`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		wg.Add(1)
		go func() {
			defer wg.Done()
			e.Request.Visit(link)
		}()
	})

	c.OnHTML("style", func(e *colly.HTMLElement) {
		mu.Lock()
		fonts.ExtractFonts(e.Text, result.FontSet)
		mu.Unlock()
	})

	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Headers.Get("Content-Type"), "text/css") {
			mu.Lock()
			fonts.ExtractFonts(string(r.Body), result.FontSet)
			result.ContainsMediaQueries = result.ContainsMediaQueries || responsive.ContainsMediaQueries(string(r.Body))
			mu.Unlock()
		} else {
			mu.Lock()
			result.VisitedUrls[r.Request.URL.String()] = true
			mu.Unlock()
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode != http.StatusOK {
			mu.Lock()
			result.BrokenLinks[r.Request.URL.String()] = true
			mu.Unlock()
		}
	})

	c.OnHTML("img", func(e *colly.HTMLElement) {
		alt := e.Attr("alt")
		src := e.Request.AbsoluteURL(e.Attr("src"))
		mu.Lock()
		result.AltTexts[alt] = append(result.AltTexts[alt], src)
		mu.Unlock()
	})

	c.OnHTML("p, h1, h2, h3, h4, h5, h6, li, span, a", func(e *colly.HTMLElement) {
		text := e.Text
		mu.Lock()
		spelling.ExtractAndCheckText(text, sc, result.MisspelledWords)
		mu.Unlock()
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

	c.OnHTML(`meta[name="viewport"]`, func(e *colly.HTMLElement) {
		mu.Lock()
		result.HasViewport[e.Request.URL.String()] = true
		mu.Unlock()
	})
}

// StartScraping begins the scraping process on the specified URL
func StartScraping(c *colly.Collector, sc *spellcheck.Trie, url string) ScrapeResult {
	result := ScrapeResult{
		FontSet:         make(map[string]bool),
		AltTexts:        make(map[string][]string),
		MisspelledWords: make(map[string]bool),
		BrokenLinks:     make(map[string]bool),
		HasViewport:     make(map[string]bool),
		VisitedUrls:     make(map[string]bool),
	}
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	SetupHandlers(c, sc, wg, &result, mu)
	c.Visit(url + "/")
	c.Wait()  // Wait for all collectors to complete, including async visits
	wg.Wait() // Wait for all goroutines to finish
	return result
}
