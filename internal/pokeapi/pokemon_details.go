package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemonDetails(url string) (Pokemon, error) {
	data, err := makeCachedRequest(c, url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("cached request failed: %w", err)
	}

	pokemonData := Pokemon{}

	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, fmt.Errorf("response decoding failed: %w", err)
	}

	return pokemonData, nil
}