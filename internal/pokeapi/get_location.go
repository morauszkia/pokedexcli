package pokeapi

import "encoding/json"

func (c *Client) ExploreLocation(locationName string) (LocationDetails, error) {
	fullURL := c.baseURL + "/location-area/" + locationName

	data, err := c.retrieveOrFetch(fullURL) 
	if err != nil {
		return LocationDetails{}, err
	}

	var locationDetails LocationDetails
	if err := json.Unmarshal(data, &locationDetails); err != nil {
		return LocationDetails{}, err
	}
	return locationDetails, nil
}