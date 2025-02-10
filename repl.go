package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pl1000100/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	pokeapiClient    pokeapi.ApiClient
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		if len(cleanText) == 0 {
			continue
		}
		args := []string{}
		if len(cleanText) > 1 {
			args = cleanText[1:]
		}

		command, exists := getCommands(cfg, &args)[cleanText[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func getCommands(cfg *config, args *[]string) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(cfg, args) },
		},
		"map": {
			name:        "map",
			description: "Displays names of next 20 location areas in the Pokemon world",
			callback:    func() error { return commandMap(cfg) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 location areas in the Pokemon world",
			callback:    func() error { return commandMapb(cfg) },
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemons in given location area",
			callback:    func() error { return commandExplore(cfg, args) },
		},
	}

}
