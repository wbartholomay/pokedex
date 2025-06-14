package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/wbartholomay/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL string
	prevLocationsURL string
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	substrings := strings.Fields(text)
	return substrings
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		t := scanner.Text()
		input := cleanInput(t)
		if len(input) == 0 { continue }

		cmd, ok := getCommands()[input[0]]
		if !ok{
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, input[1:]...)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}