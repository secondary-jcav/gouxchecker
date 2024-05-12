package images

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAltText(t *testing.T) {
	altTexts := map[string][]string{
		"":                  {"https://website.com/static/ccp.svg", "https://website.com/static/arrow.png"},
		"duplicate text":    {"https://website.com/static/aws.svg", "https://website.com/static/angular.svg"},
		"another duplicate": {"https://website.com/static/gcp.svg", "https://website.com/static/azure.svg"},
		"test":              {"https://website.com/static/test.svg"},
	}

	noAlts, duplicateAlts := CheckAltText(altTexts)
	expectedNoAlts := []string{"https://website.com/static/ccp.svg", "https://website.com/static/arrow.png"}
	expectedDuplicateAlts := []string{"https://website.com/static/aws.svg", "https://website.com/static/angular.svg", "https://website.com/static/gcp.svg", "https://website.com/static/azure.svg"}

	assert := assert.New(t)
	assert.Equal(len(noAlts), 2, "Images without alt text not detected correctly")
	assert.Equal(len(duplicateAlts), 4, "Images with duplicate alt text not detected correctly")

	for _, url := range expectedNoAlts {
		assert.Contains(noAlts, url, "missing element in noAlts: %s", url)
	}

	for _, url := range expectedDuplicateAlts {
		assert.Contains(duplicateAlts, url, "missing element in duplicateAlts: %s", url)
	}

}
