package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nmap: Displays the map\nexplore: Explore local area for pokemon\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}

func commandMap(cfg *Config, args ...string) error {
	var pageURL *string
	if cfg.Next != "" {
		pageURL = &cfg.Next
	}

	resp, err := cfg.pokeapiClient.FetchLocations(pageURL)
	if err != nil {
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

func commandMapb(cfg *Config, args ...string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	pageURL := &cfg.Previous
	resp, err := cfg.pokeapiClient.FetchLocations(pageURL)
	if err != nil {
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
