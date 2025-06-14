package main

import (
	"errors"
	"fmt"
)

func CommandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("must provide a location. Example usage: explore <location-name>")
	}
	
	url := "https://pokeapi.co/api/v2/location-area/" + params[0]
	pokemonRes, err := cfg.pokeapiClient.ExploreLocation(url)
	if err != nil {
		return fmt.Errorf("call to PokeAPI failed: %w", err)
	}

	for _, pokemon := range pokemonRes.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}