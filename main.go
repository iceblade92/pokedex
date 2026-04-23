package main

import (
	"time"

	"github.com/iceblade92/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	Repl(cfg)

}
