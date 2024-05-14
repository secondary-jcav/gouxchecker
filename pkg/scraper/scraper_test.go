package scraper

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/secondary-jcav/gouxchecker/pkg/spelling"

	"github.com/gocolly/colly/v2"

	"github.com/stretchr/testify/assert"
)

func TestInitializeCollector(t *testing.T) {
	domain := "https://example.com"
	collector := InitializeCollector(domain)
	assert.NotNil(t, collector, "Collector should not be nil")
	assert.Contains(t, collector.AllowedDomains, "example.com", "Domain should be in the list of allowed domains")
}

func TestStartSCraping(t *testing.T) {

	sc := spelling.InitSpellchecker()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			htmlContent := `
<!DOCTYPE html>
<html>
	<head>
		<style>
			body {
				font-family: Arial, sans-serif;
			}
		</style>
	</head>
	<body>
		<h1>Welcomeee!</h1>
		<img src="example.jpg" alt="Descriptive text about the image">
	</body>
</html>`
			fmt.Fprintf(w, htmlContent)
		})
		fmt.Println("Server starting on port 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	time.Sleep(1 * time.Second)
	c := colly.NewCollector()

	fontSet, altText, misspelledWords, _ := StartScraping(c, sc, "http://localhost:8080")

	assert := assert.New(t)
	assert.NotEmpty(fontSet)
	assert.NotEmpty(altText)
	assert.NotEmpty(misspelledWords)

}
