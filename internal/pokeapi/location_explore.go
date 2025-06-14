package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ExploreLocation(url string) (RespDetailedLocations, error) {
	data, err := makeCachedRequest(c, url)
	if err != nil {
		return RespDetailedLocations{}, fmt.Errorf("cached request failed: %w", err)
	}

	pokemonData := RespDetailedLocations{}

	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return RespDetailedLocations{}, fmt.Errorf("response decoding failed: %w", err)
	}

	return pokemonData, nil
}