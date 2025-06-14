package main

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("must provide a pokemon. Example usage: inspect <pokemon-name>")
	}

	pokemonName := params[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
	}

	return nil

}