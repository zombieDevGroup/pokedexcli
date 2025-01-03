package main

import (
	"strings"
)

func cleanInput(text string) []string {
	if len(strings.TrimSpace(text)) == 0 {
		return []string{}
	}

	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)

	}

	return words
}
