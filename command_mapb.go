package main

import (
	"fmt"
)

func commandMapb(c *Config, args ...string) error {
	var locationAreasURL string
	if c.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		locationAreasURL = *c.prev
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