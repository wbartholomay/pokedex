package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	substrings := strings.Fields(text)
	return substrings
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config {
		Next : "",
		Prev : "",
	}

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

		err := cmd.callback(&cfg)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}