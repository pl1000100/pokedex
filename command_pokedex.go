package main

import "fmt"

func commandPokedex(cfg *config) error {
	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
