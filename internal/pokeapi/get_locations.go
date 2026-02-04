package pokeapi

import "encoding/json"

func (c *Client) GetLocationAreas(url string) (LocationAreasList, error) {
	var fullURL string
	if url == "" {
		fullURL = c.baseURL + "/location-area"
	} else {
		fullURL = url
	}

	data, err := c.retrieveOrFetch(fullURL)
	if err != nil {
		return LocationAreasList{}, err
	}

	var locationAreasList LocationAreasList
	if err := json.Unmarshal(data, &locationAreasList); err != nil {
		return LocationAreasList{}, err
	}
	return locationAreasList, nil
}