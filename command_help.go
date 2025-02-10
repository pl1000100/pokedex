package main

import (
	"fmt"

	"github.com/pl1000100/pokedex/internal/pokeapi"
)

func commandHelp(client *pokeapi.ApiClient) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range getCommands(client) {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
