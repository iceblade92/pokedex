package main

import (
	"fmt"
	"os"
)

func commandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nmap: Displays the map\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandMap(cfg *Config) error {
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

func commandMapb(cfg *Config) error {
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
