package pokeapi

import (
	"net/http"
	"time"

	"github.com/morauszkia/pokedexcli/internal/pokecache"
)

type Client struct {
	client 		*http.Client
	cache		*pokecache.Cache
	baseURL		string	
}

func NewClient (cacheInterval time.Duration) *Client {
	client := Client{
		client: &http.Client{},
		cache: pokecache.NewCache(cacheInterval),
		baseURL: "https://pokeapi.co/api/v2",
	}
	return &client
}