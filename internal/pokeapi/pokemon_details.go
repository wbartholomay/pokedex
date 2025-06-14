package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemonDetails(url string) (RespPokemonDetails, error) {
	fmt.Println(url)
	data, err := makeCachedRequest(c, url)
	if err != nil {
		return RespPokemonDetails{}, fmt.Errorf("cached request failed: %w", err)
	}

	pokemonData := RespPokemonDetails{}

	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return RespPokemonDetails{}, fmt.Errorf("response decoding failed: %w", err)
	}

	return pokemonData, nil
}