package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	url := cfg.nextLocationsURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	return commonMap(cfg, url)
}

func commandMapb(cfg *config) error {
	url := cfg.prevLocationsURL
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	return commonMap(cfg, url)
}

func commonMap(cfg *config, url string) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return fmt.Errorf("call to PokeAPI failed: %w", err)
	}


	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}