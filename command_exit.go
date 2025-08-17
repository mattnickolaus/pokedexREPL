package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex...")
	fmt.Println(Pikachu)
	os.Exit(0)
	return nil
}
