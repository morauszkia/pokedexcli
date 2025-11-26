package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startREPL() {
	s := bufio.NewScanner(os.Stdin)
	config := NewConfig()

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
		
		err := command.callback(&config)
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