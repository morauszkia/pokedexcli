package main

import (
	"bufio"
	"fmt"
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandStr]
		if !exists  {
			fmt.Print("Unknown command! Type 'help' to list available commands.\n")
			continue
		}
		
		err := command.callback(&config, args...)
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