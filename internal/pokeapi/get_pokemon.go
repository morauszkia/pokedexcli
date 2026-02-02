package pokeapi

import "encoding/json"

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	fullURL := c.baseURL + "/pokemon/" + pokemonName

	jsonData, err := c.retrieveOrFetch(fullURL)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(jsonData, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}