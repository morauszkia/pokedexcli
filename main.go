package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		words := cleanInput(input)
		commandStr := words[0]
		command := commands[commandStr]
		if command.callback == nil  {
			fmt.Print("Unknown command\n")
			continue
		}
		command.callback()
	}
}

