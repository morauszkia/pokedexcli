package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: inspect <pokemon-name>")
	}

	pokemonName := args[0]
	pokemon, ok := c.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You haven't caught %s", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("- %s\n", typeInfo.Type.Name)
	}
	
	return nil
}