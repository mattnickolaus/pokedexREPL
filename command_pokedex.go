package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println()

	if len(cfg.Pokedex) == 0 {
		return errors.New("No pokemon have been caught")
	}

	for k, _ := range cfg.Pokedex {
		fmt.Printf(" - %s\n", k)
	}

	fmt.Println()

	return nil
}
