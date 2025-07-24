package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mattnickolaus/pokedexREPL/internal/pokeapi"
)

var commandRegister map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func initCommandRegister() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the nearest(next) location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous location-areas",
			callback:    commandMapBack,
		},
	}
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commandRegister = initCommandRegister()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		if _, ok := commandRegister[command]; ok {
			c := commandRegister[command]
			err := c.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
