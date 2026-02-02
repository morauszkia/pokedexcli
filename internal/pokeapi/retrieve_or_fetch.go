package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) retrieveOrFetch(fullURL string) ([]byte, error) {
	cachedData, ok := c.cache.Get(fullURL)
	if ok {
		return cachedData, nil
	}

	res, err := c.client.Get(fullURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(fullURL, data)
	return data, nil
}