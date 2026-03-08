package responsive

import (
	"sort"
	"strings"
)

func ContainsMediaQueries(cssContent string) bool {
	return strings.Contains(cssContent, "media")
}

// CheckResponsiveUrls returns visited URLs that lack a <meta name="viewport"> tag.
func CheckResponsiveUrls(visitedUrls, hasViewport map[string]bool) []string {
	var results []string
	for url := range visitedUrls {
		if !hasViewport[url] {
			results = append(results, url)
		}
	}
	sort.Strings(results)
	return results
}
