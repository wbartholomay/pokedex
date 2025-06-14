package main

import(
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
}
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}