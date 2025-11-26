package main

import (
	"fmt"
)

func commandMap(c *Config) error {
	var locationAreasURL string
	if c.next == nil {
		locationAreasURL = ""
	} else {
		locationAreasURL = *c.next
	}

	data, err := c.client.GetLocationAreas(locationAreasURL)
	if err != nil {
		return err
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	c.next = data.Next
	c.prev = data.Previous

	return nil
}