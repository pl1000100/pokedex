package main

import (
	"time"

	"github.com/pl1000100/pokedex/internal/pokeapi"
)

func main() {
	// TODO: move client to config
	client := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: client,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(config)
}
