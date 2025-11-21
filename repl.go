package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startREPL() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()

		words := cleanInput(s.Text())
		if len(words) == 0 {
			fmt.Println("Please enter command")
			continue
		}

		commandStr := words[0]
		command := getCommands()[commandStr]
		if command.callback == nil  {
			fmt.Print("Unknown command! Type 'help' to list available commands.\n")
			continue
		}
		
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:			"help",
			description: 	"Lists available commands",
			callback: 		commandHelp,
		},
		"exit": {
			name: 			"exit",
			description:	"Exits the Pokedex",
			callback:		commandExit,
		},
	}
}