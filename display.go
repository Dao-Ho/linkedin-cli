package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func displayFeed() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromString("LINKEDIN CLI"),
	).Render()
	fmt.Println("[Placeholder: Feed will load here]")
}
