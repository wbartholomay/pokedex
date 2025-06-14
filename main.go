package main

import (
	"github.com/wbartholomay/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}