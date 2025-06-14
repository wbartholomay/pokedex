package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func CommandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("must provide a pokemon. Example usage: catch <pokemon-name>")
	}

	pokemonName := params[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pokemonRes, err := cfg.pokeapiClient.GetPokemonDetails(url)
	if err != nil {
		return err
	}

	exp := pokemonRes.BaseExperience
	
	if rand.IntN(609) > exp / 2 {
		fmt.Printf("%s was caught!\n", pokemonName)
		pokedex[pokemonName] = Pokemon{
			Name: pokemonRes.Name,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil

}