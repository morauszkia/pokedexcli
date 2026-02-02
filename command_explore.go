package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: explore <location-name>")
	}

	locationName := args[0]
	location, err := c.client.ExploreLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	if len(location.PokemonEncounters) == 0 {
		fmt.Printf("No pokemon encountered...\n")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}