package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	client 		*http.Client
	baseURL		string	
}

func (c *Client) GetLocationAreas(url string) (LocationAreasList, error) {
	var fullURL string
	if url == "" {
		fullURL = c.baseURL + "/location-area"
	} else {
		fullURL = url
	}
	
	res, err := c.client.Get(fullURL)
	if err != nil {
		return LocationAreasList{}, err
	}
	if res.StatusCode != http.StatusOK {
		return LocationAreasList{}, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasList{}, err
	}

	var locationAreasList LocationAreasList
	if err := json.Unmarshal(data, &locationAreasList); err != nil {
		return LocationAreasList{}, err
	}
	return locationAreasList, nil
}

func NewClient () *Client {
	client := Client{
		client: &http.Client{},
		baseURL: "https://pokeapi.co/api/v2",
	}
	return &client
}