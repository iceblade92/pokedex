package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	text := `Welcome to the Pokedex!
Usage:


map: Displays the map
explore: Explore local area for pokemon
help: Displays a help message
exit: Exit the Pokedex`
	fmt.Println(text)
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

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := rand.Intn(400)
	if chance > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)

	}
	fmt.Println("Types:")
	for _, element := range pokemon.Types {
		fmt.Printf("  - %s\n", element.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have no pokemon try catching first!")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
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
