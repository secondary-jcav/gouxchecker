package responsive

import "strings"

func ContainsMediaQueries(cssContent string) bool {
	return strings.Contains(cssContent, "media")

}

func CheckResponsiveUrls(visitedUrls, nonNativeUrls map[string]bool) []string {
	// add to result if
	//a) visitedUrl does not exist in nonNative or
	//b) it exists and the value is true
	var results []string
	for url := range visitedUrls {
		if nonNative, exists := nonNativeUrls[url]; !exists || nonNative {
			results = append(results, url)
		}
	}
	return results
}
