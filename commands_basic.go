package main

import(
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(cfg *config, params ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
	"exit" : cliCommand{
		name : "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help" : cliCommand{
		name : "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map" : cliCommand{
		name : "map",
		description: "Displays the next 20 locations",
		callback: commandMap,
	},
	"mapb" : cliCommand{
		name : "mapb",
		description: "Displays the previous 20 locations",
		callback: commandMapb,
	},
	"explore" : cliCommand{
		name : "explore",
		description: "Takes one argument <location-name>. Displays all of the Pokemon available in the given location.",
		callback: CommandExplore,
	},
	"catch" : cliCommand{
		name : "catch",
		description: "Takes one argument <pokemon-name>. Attempts to catch the Pokemon, and returns whether or not the catch succeeded.",
		callback: CommandCatch,
	},
}
}


func commandExit(cfg *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}