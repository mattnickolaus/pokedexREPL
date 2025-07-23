package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandRegister map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startRepl() {
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
			err := c.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
