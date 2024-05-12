package fonts

import (
	"testing"
)

func TestExtractFonts(t *testing.T) {
	content := `
	h1 {
		color: var(--green);
		font-size: 48px;
		font-family: "Roboto";
		font-weight: 700;
		height: 63px;
	  }
	  
	  h2 {
		color: var(--light-grey);
		font-size: 28px;
		font-family: "Arial";
		font-weight: 300;
		height: fit-content;
		width: 60vw;
	  }
	  
	  .hero-button {
		background-color: transparent;
		border: none;
		color: var(--green);
		font-size: 24px;
		font-family: "Helvetica";
		font-weight: 500;
	  }
    `

	// Expected fonts in the fontSet map
	expectedFonts := map[string]bool{
		"Arial":     true,
		"Roboto":    true,
		"Helvetica": true,
	}

	fontSet := make(map[string]bool)
	ExtractFonts(content, fontSet)

	for font := range expectedFonts {
		if !fontSet[font] {
			t.Errorf("Font '%s' not found in fontSet", font)
		}
	}

	for font := range fontSet {
		if !expectedFonts[font] {
			t.Errorf("Unexpected font '%s' found in fontSet", font)
		}
	}
}
