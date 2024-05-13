package output

import (
	"fmt"
	"os"
	"sort"
)

func ImageAlts(noAlts []string, duplicateAlts []string) {

	if len(noAlts) > 0 {
		fmt.Println("Found images without alt text")
		noAltsFile, err := os.Create("results/images_no_alt.txt")
		if err != nil {
			fmt.Println("Error when creating the file:", err)
			return
		}
		defer noAltsFile.Close()
		for _, alt := range noAlts {
			fmt.Fprintln(noAltsFile, alt)
		}

	} else {
		fmt.Println("Didn't find images without alt text")
	}

	if len(duplicateAlts) > 0 {
		fmt.Println("Found images with duplicate alt text")
		duplicateAltsFile, err := os.Create("results/images_duplicate_alt.txt")
		if err != nil {
			fmt.Println("Error when creating the file:", err)
			return
		}
		defer duplicateAltsFile.Close()
		for _, alt := range duplicateAlts {
			fmt.Fprintln(duplicateAltsFile, alt)
		}
	} else {
		fmt.Println("Didn't find images with duplicate alt text")
	}

}

func FontsResults(fontSet map[string]bool) {
	fontsFile, err := os.Create("results/fonts.txt")
	if err != nil {
		fmt.Println("Error when creating the file:", err)
		return
	}
	if len(fontSet) > 3 {

		fmt.Println("Too many fonts")

	} else {
		fmt.Println("3 or fewer fonts in use: OK")
	}
	for font := range fontSet {
		fmt.Fprintln(fontsFile, font)
	}
}

func Typos(misspelledWords map[string]bool) {
	if len(misspelledWords) != 0 {
		var keys []string
		for key := range misspelledWords {
			keys = append(keys, key)
		}
		fmt.Println("Possible typos found")
		sort.Strings(keys)
		file, err := os.Create("results/typos.txt")
		if err != nil {
			fmt.Println("Error when creating the file:", err)
			return
		}
		defer file.Close()
		for _, key := range keys {
			fmt.Fprintln(file, key)
		}
	}

}

func BrokenLinks(brokenLinks map[string]bool) {
	if len(brokenLinks) > 0 {
		fmt.Println("Found broken links")
		file, err := os.Create("results/broken_links.txt")
		if err != nil {
			fmt.Println("Error when creating the file:", err)
			return
		}
		defer file.Close()
		for link := range brokenLinks {
			fmt.Fprintln(file, link)
		}

	} else {
		fmt.Println("No broken links found")
	}
}
