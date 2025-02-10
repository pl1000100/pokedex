package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pl1000100/pokedex/internal/pokeapi"
	"github.com/pl1000100/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	const interval = 5 * time.Second

	client := pokeapi.ApiClient{
		Cache: pokecache.NewCache(interval),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean_text := cleanInput(scanner.Text())
		if len(clean_text) == 0 {
			continue
		}
		command, exists := getCommands(&client)[clean_text[0]]
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

func getCommands(client *pokeapi.ApiClient) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(client) },
		},
		"map": {
			name:        "map",
			description: "Displays names of next 20 location areas in the Pokemon world",
			callback:    func() error { return commandMap(client) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 location areas in the Pokemon world",
			callback:    func() error { return commandMapb(client) },
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemons in given location area",
			callback:    func() error { return commandExplore(client) },
		},
	}

}
