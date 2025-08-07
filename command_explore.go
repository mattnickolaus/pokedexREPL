package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {
	fmt.Println()

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)
	locDetails, err := cfg.pokeapiClient.GetLocationDetails(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, e := range locDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", e.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
