package main

import (
	"time"

	"github.com/morauszkia/pokedexcli/internal/pokeapi"
)

type Config struct {
	prev		*string
	next		*string
	client		*pokeapi.Client
}

func NewConfig() Config {
	return Config{
		prev: nil,
		next: nil,
		client: pokeapi.NewClient(time.Duration(1 * float64(time.Second))),
	}
}