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

var commands map[string]cliCommand

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, commandStruct := range commands {
		fmt.Printf("%s: %s\n", commandStruct.name, commandStruct.description)
	}
	return nil
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func init() {
	commands = make(map[string]cliCommand)
	commands["help"] = cliCommand{
		name:			"help",
		description: 	"Lists available commands",
		callback: 		commandHelp,
	}
	commands["exit"] = cliCommand{
		name: 			"exit",
		description:	"Exits the Pokedex",
		callback:		commandExit,
	}
}