package main

import "fmt"

func commandExplore(cfg *config, args *[]string) error {
	if *args == nil || len(*args) < 1 {
		fmt.Println("Missing location")
		return nil
	}
	locName := (*args)[0]
	loc, err := cfg.pokeapiClient.SingleLocationArea(locName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", locName)
	for _, p := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
