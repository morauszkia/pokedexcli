package main

import (
	"fmt"
	"os"
	"strings"
)

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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name: 			"exit",
		description:	"Exits the Pokedex",
		callback:		commandExit,
	},
}