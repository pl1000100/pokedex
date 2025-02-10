package main

import (
	"github.com/pl1000100/pokedex/internal/pokeapi"
)

// TODO: move GET call to single file, handling cache and not in this two functions
func commandMap(client *pokeapi.ApiClient) error {
	pokeapi.List_location_area(client.Next, client)
	return nil
}

func commandMapb(client *pokeapi.ApiClient) error {
	pokeapi.List_location_area(client.Previous, client)
	return nil
}
