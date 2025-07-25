package main

import (
	"fmt"
	"time"

	"github.com/mattnickolaus/pokedexREPL/internal/pokeapi"
)

func main() {
	fmt.Printf("%v\n\n", PokedexWelcome)

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
