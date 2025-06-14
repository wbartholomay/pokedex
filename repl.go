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

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		t := scanner.Text()
		input := cleanInput(t)
		if len(input) == 0 { continue }

		fmt.Println("Your command was: " + input[0])
	}
}