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
		fmt.Printf("Your command was: %s\n", commandStr)
	}
}

