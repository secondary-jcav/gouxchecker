package fonts

import (
	"regexp"
	"strings"
)

func ExtractFonts(content string, fontSet map[string]bool) {
	fontRe := regexp.MustCompile(`font-family:([^;}]+)`)
	matches := fontRe.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			font := strings.TrimSpace(strings.ReplaceAll(match[1], "\"", ""))
			fontSet[font] = true
		}
	}
}
