package main

import (
	"errors"
	"fmt"

	"github.com/morauszkia/pokedexcli/internal/pokeapi"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Usage: catch <pokemon-name>")
	}

	pokemonName := args[0]
	pokemon, err := c.client.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if isCatchSuccessful(pokemon) {
		fmt.Println("Success!")
	}

	return nil
}


func isCatchSuccessful(pokemon pokeapi.Pokemon) bool {
	return false
}