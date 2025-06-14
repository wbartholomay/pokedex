package main

import (
	"fmt"
)

func CommandPokedex(cfg *config, params ...string) error {
	fmt.Println("Your Pokedex:")

	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}