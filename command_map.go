package main

import (
	"fmt"
)

// TODO: move GET call to single file, handling cache and not in this two functions
func commandMap(cfg *config) error {
	data, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	for _, result := range data.Results {
		fmt.Println(result.Name)
	}
	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous
	return nil
}

func commandMapb(cfg *config) error {
	data, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	for _, result := range data.Results {
		fmt.Println(result.Name)
	}
	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous
	return nil
}
