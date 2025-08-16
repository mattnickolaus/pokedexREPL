package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	fmt.Println()

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Change this to be getPokemon(pokemon)
	pokemonDetails, err := cfg.pokeapiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}

	return nil
}

func isCaught(baseExp int) (bool, error) {
	catchRate := 10.0 / baseExp
	return false, nil
}
