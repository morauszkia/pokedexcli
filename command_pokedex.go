package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *Config, args ...string) error {
	if len(args) > 0 {
		return errors.New("Too many arguments. Usage: pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}