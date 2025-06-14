package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationResContent struct {
	Count int
	Next string
	Previous string
	Results []location
}

type location struct {
	Name string
	Url string
}

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
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("call to PokeAPI failed: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed. Status Code: %v     Status Description: %v", res.StatusCode, res.Status)
	}
	
	var jsonData locationResContent
	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&jsonData)
	if err != nil {
		return fmt.Errorf("JSON decoding failed: %w", err)
	}

	cfg.nextLocationsURL = jsonData.Next
	cfg.prevLocationsURL = jsonData.Previous

	for _, loc := range jsonData.Results {
		fmt.Println(loc.Name)
	}

	return nil
}