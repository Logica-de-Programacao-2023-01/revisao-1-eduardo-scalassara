package main

import (
	"errors"
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(text) == 0 || strings.TrimSpace(text) == "" {
		return 0, errors.New("texto vazio")
	}

	words := strings.Fields(text)
	wordCount := len(words)
	letterCount := 0

	for _, word := range words {
		letterCount += len(word)
	}

	average := float64(letterCount) / float64(wordCount)
	return average, nil
}


