package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

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
		c.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}


func isCatchSuccessful(pokemon pokeapi.Pokemon) bool {
	chance := math.Max(0.1, math.Min(0.8, 0.7 - 0.002 * float64(pokemon.BaseExperience)))
	return rand.Float64() > chance
}