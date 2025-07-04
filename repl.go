package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Trev-D-Dev/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()

			cInput := cleanInput(input)

			commName := cInput[0]

			comm, ok := getCommands()[commName]

			if ok {
				var err error
				if commName == "explore" {
					err = comm.callback(cfg, cInput[1])
				} else {
					err = comm.callback(cfg, "")
				}

				if err != nil {
					fmt.Printf("error: %v", err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints usage and commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 pokemon world area location names",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 pokemon world area location names",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon that can be found at a specific world area location",
			callback:    commandExplore,
		},
	}
}
