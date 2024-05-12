package spelling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractAndCheckText(t *testing.T) {

	test := "This Zentence should have 2 typ0z"
	sc := InitSpellchecker()
	misspelledWords := make(map[string]bool)

	expectedMisspelled := map[string]bool{
		"zentence": true,
		"typ0z":    true,
	}

	ExtractAndCheckText(test, sc, misspelledWords)

	assert.Equal(t, expectedMisspelled, misspelledWords, "The misspelled words map does not match the expected output")

}
