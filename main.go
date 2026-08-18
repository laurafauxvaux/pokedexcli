package main

import (
	"time"

	"github.com/laurafauxvaux/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	newPokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:       newPokedex,
	}
	repl(cfg)
}
