package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (LocationDetails, error) {
	fullURL := c.baseURL + "/location-area/" + locationName

	cachedData, ok := c.cache.Get(fullURL)
	if ok {
		var locationDetails LocationDetails
		if err := json.Unmarshal(cachedData, &locationDetails); err != nil {
			return LocationDetails{}, err
		}
		return locationDetails, nil
	}

	res, err := c.client.Get(fullURL)
	if err != nil {
		return LocationDetails{}, err
	}
	if res.StatusCode != http.StatusOK {
		return LocationDetails{}, fmt.Errorf("Unexpected status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	var locationDetails LocationDetails
	if err := json.Unmarshal(data, &locationDetails); err != nil {
		return LocationDetails{}, err
	}
	c.cache.Add(fullURL, data)
	return locationDetails, nil
}