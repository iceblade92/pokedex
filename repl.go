package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iceblade92/pokedex/internal/pokeapi"
)

type Config struct {
	pokeapiClient *pokeapi.Client
	Next          string
	Previous      string
}

func Repl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		str := cleanInput(text)
		if len(str) == 0 {
			fmt.Println("No text prompt provided")
			continue
		}
		args := []string{}
		if len(str) > 1 {
			args = str[1:]
		}

		command, exists := getCommands()[str[0]]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func cleanInput(text string) []string {
	str := text
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	result := strings.Fields(str)
	return result
}

func getCommands() map[string]cliCommand {
	command := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help usage for Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Map shows current 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Map back shows previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows presant pokemon in local area",
			callback:    commandExplore,
		},
	}
	return command
}
