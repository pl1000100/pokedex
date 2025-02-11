package main

import "fmt"

func commandInspect(cfg *config, args *[]string) error {
	if args == nil || len(*args) < 1 {
		fmt.Println("You need to specify name!")
		return nil
	}
	pokemonName := (*args)[0]
	poke, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}
