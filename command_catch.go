package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args *[]string) error {
	if args == nil || len(*args) < 1 {
		fmt.Println("Missing Pokemon name")
		return nil
	}
	pokemonName := (*args)[0]
	pokemon, err := cfg.pokeapiClient.SinglePokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if rand.IntN(100) > catchChance(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.pokedex[pokemonName] = pokemon
	return nil
}

func catchChance(exp int) int {
	baseChance := 20
	return baseChance + int(1000/exp)
}
