package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const api = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
}

type locationAreaResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
	}
	return command
}

func cleanInput(text string) []string {
	str := text
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	result := strings.Fields(str)
	return result
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		str := cleanInput(text)
		if len(str) == 0 {
			fmt.Println("No text prompt provided")
			continue
		}
		command, exists := getCommands()[str[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func commandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nmap: Displays the map\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func fetchLocations(url string, cfg *Config) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resp locationAreaResp
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	if resp.Next != nil {
		cfg.Next = *resp.Next
	} else {
		cfg.Next = ""
	}
	if resp.Previous != nil {
		cfg.Previous = *resp.Previous
	} else {
		cfg.Previous = ""
	}
	return nil
}

func commandMap(cfg *Config) error {
	url := api
	if cfg.Next != "" {
		url = cfg.Next
	}
	return fetchLocations(url, cfg)
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return fetchLocations(cfg.Previous, cfg)
}
