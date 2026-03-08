package responsive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsMediaQueries(t *testing.T) {
	assert := assert.New(t)

	cssWithMedia := "@media (max-width: 600px) { .class { display: none; } }"
	assert.True(ContainsMediaQueries(cssWithMedia), "CSS with media queries should return true")

	cssWithoutMedia := ".class { display: block; }"
	assert.False(ContainsMediaQueries(cssWithoutMedia), "CSS without media queries should return false")

	emptyCSS := ""
	assert.False(ContainsMediaQueries(emptyCSS), "Empty CSS should return false")
}

func TestCheckResponsiveUrls(t *testing.T) {
	assert := assert.New(t)

	visitedUrls := map[string]bool{
		"http://example.com": true,
		"http://another.com": true,
		"http://sample.com":  true,
	}

	// example.com has viewport → responsive, the others do not
	hasViewport := map[string]bool{
		"http://example.com": true,
	}

	results := CheckResponsiveUrls(visitedUrls, hasViewport)
	expectedResults := []string{"http://another.com", "http://sample.com"}
	assert.Equal(expectedResults, results, "Should return non-responsive URLs, sorted")

	emptyVisited := make(map[string]bool)
	emptyHasViewport := make(map[string]bool)
	resultsEmpty := CheckResponsiveUrls(emptyVisited, emptyHasViewport)
	assert.Empty(resultsEmpty, "Should return an empty slice for empty input maps")
}
