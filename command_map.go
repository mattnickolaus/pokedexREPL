package main

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	fmt.Println()

	areas, err := cfg.pokeapiClient.GetLocationArea(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	for _, v := range areas.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}

func commandMapBack(cfg *Config) error {
	fmt.Println()

	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	areas, err := cfg.pokeapiClient.GetLocationArea(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	for _, v := range areas.Results {
		fmt.Println(v.Name)
	}
	fmt.Println()

	return nil
}
