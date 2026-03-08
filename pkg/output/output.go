package output

import (
	"fmt"
	"os"
	"sort"
)

func writeLinesToFile(path string, lines []string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Error when creating the file:", err)
		return
	}
	defer f.Close()
	for _, line := range lines {
		fmt.Fprintln(f, line)
	}
}

func ImageAlts(noAlts []string, duplicateAlts []string) {
	if len(noAlts) > 0 {
		fmt.Println("Found images without alt text")
		writeLinesToFile("results/images_no_alt.txt", noAlts)
	} else {
		fmt.Println("Didn't find images without alt text")
	}

	if len(duplicateAlts) > 0 {
		fmt.Println("Found images with duplicate alt text")
		writeLinesToFile("results/images_duplicate_alt.txt", duplicateAlts)
	} else {
		fmt.Println("Didn't find images with duplicate alt text")
	}
}

func FontsResults(fontSet map[string]bool) {
	fonts := make([]string, 0, len(fontSet))
	for font := range fontSet {
		fonts = append(fonts, font)
	}
	writeLinesToFile("results/fonts.txt", fonts)
	if len(fontSet) > 3 {
		fmt.Println("Too many fonts")
	} else {
		fmt.Println("3 or fewer fonts in use: OK")
	}
}

func Typos(misspelledWords map[string]bool) {
	if len(misspelledWords) == 0 {
		return
	}
	keys := make([]string, 0, len(misspelledWords))
	for key := range misspelledWords {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println("Possible typos found")
	writeLinesToFile("results/typos.txt", keys)
}

func BrokenLinks(brokenLinks map[string]bool) {
	if len(brokenLinks) == 0 {
		fmt.Println("No broken links found")
		return
	}
	fmt.Println("Found broken links")
	links := make([]string, 0, len(brokenLinks))
	for link := range brokenLinks {
		links = append(links, link)
	}
	writeLinesToFile("results/broken_links.txt", links)
}

func NonResponsivePages(containsMediaQueries bool, nonResponsivePages []string) {
	if len(nonResponsivePages) > 0 {
		fmt.Println("Pages missing <meta name=\"viewport\"> (likely not mobile-responsive)")
		if !containsMediaQueries {
			fmt.Println("No CSS media queries detected either")
		}
		lines := append([]string{"Pages missing viewport meta tag:"}, nonResponsivePages...)
		writeLinesToFile("results/responsiveness.txt", lines)
	} else {
		fmt.Println("All pages have viewport meta tag: OK")
	}
}
