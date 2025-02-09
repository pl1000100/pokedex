package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pl1000100/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type locationAreaResponse struct {
	Count    uint
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}

type Config struct {
	next     *string
	previous *string
	cache    *(pokecache.Cache)
}

func startRepl() {
	const interval = 5 * time.Second
	first_locations := "https://pokeapi.co/api/v2/location-area"
	cache := pokecache.NewCache(interval)
	config := Config{next: &first_locations,
		cache: cache}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean_text := cleanInput(scanner.Text())
		if len(clean_text) == 0 {
			continue
		}
		command, exists := getCommands(&config)[clean_text[0]]
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

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit(config) },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config) },
		},
		"map": {
			name:        "map",
			description: "Displays names of next 20 location areas in the Pokemon world",
			callback:    func() error { return commandMap(config) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 location areas in the Pokemon world",
			callback:    func() error { return commandMapb(config) },
		},
	}

}
