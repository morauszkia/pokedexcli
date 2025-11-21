package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Split(strings.ToLower(text), " ")
	var filteredWords []string 
	for _, word := range words {
		if word != "" {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}
