package main

import (
	"fmt"
)

func commandHelp(cfg *config, args *[]string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range getCommands(cfg, args) {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
