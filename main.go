package main

import (
	"time"

	"github.com/iceblade92/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	//cache here so they dont panic from nil values
	cfg := &Config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	Repl(cfg)

}
