package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	fmt.Println()

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]

	pokemonDetails, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("%s has not been caught", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pokemonDetails.Height)
	fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemonDetails.Stats {
		statValue := stat.BaseStat
		statName := stat.Stat.Name
		fmt.Printf("  -%s: %d\n", statName, statValue)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemonDetails.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	fmt.Println()

	return nil
}
