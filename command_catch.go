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

	caught, err := isCaught(pokemonDetails.BaseExperience)
	if err != nil {
		return err
	}
	if !caught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.Pokedex[pokemonName] = pokemonDetails
	return nil
}

func isCaught(baseExp int) (bool, error) {
	if baseExp == 0 {
		return false, fmt.Errorf("pokemon not found")
	}
	if baseExp > 400 {
		baseExp = 399
	}
	ratio := float64(baseExp) / 400
	catchRate := int(ratio * 100)
	rng := rand.Intn(100)
	// fmt.Printf("If RNG: %d > catch rate: %d\n", rng, catchRate)
	if rng > catchRate {
		return true, nil
	}
	return false, nil
}
