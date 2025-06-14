package main

import (
	"github.com/wbartholomay/pokedexcli/internal/pokeapi"
	"time"
)

var pokedex map[string]Pokemon

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	pokedex = make(map[string]Pokemon)

	startRepl(cfg)
}