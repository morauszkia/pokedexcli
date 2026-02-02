package main

import "fmt"

func commandHelp(c *Config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, commandStruct := range commands {
		fmt.Printf("%s: %s\n", commandStruct.name, commandStruct.description)
	}
	fmt.Println()
	return nil
}
