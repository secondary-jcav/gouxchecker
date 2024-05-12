package images

func CheckAltText(altTexts map[string][]string) ([]string, []string) {
	if altTexts == nil {
		return nil, nil
	}

	noAlts := []string{}
	duplicateAlts := []string{}
	altOccurrences := make(map[string]bool)

	for alt, srcs := range altTexts {
		if len(srcs) > 1 && alt != "" {
			// If there are multiple source URLs for this alt text and alt text is not empty,
			// add all source URLs to duplicateAlts slice
			duplicateAlts = append(duplicateAlts, srcs...)
		}
		for _, src := range srcs {
			if alt == "" {
				// If alt text is empty, add the source URL
				noAlts = append(noAlts, src)
			} else {
				// If alt text is not empty, mark it as occurred
				key := alt + ":" + src
				altOccurrences[key] = true
			}
		}
	}

	return noAlts, duplicateAlts
}
