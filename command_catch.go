package main

import (
	"errors"
	"fmt"
)

func CommandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("must provide a pokemon. Example usage: catch <pokemon-name>")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + params[0]
	pokemonRes, err := cfg.pokeapiClient.GetPokemonDetails(url)
	if err != nil {
		return err
	}

	exp := pokemonRes.BaseExperience
	fmt.Println(exp)

	return nil

}