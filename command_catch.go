package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no Pokemon name was given")
	}

	pokemon := args[0]

	pokemonDetails, err := cfg.pokeapiClient.PokemonDetails(pokemon)

	if err != nil {
		return err
	}

	chance := max(min(100-pokemonDetails.BaseExperience/4, 90), 10)

	attempt := rand.Intn(100)

	fmt.Println("Throwing a Pokeball at", pokemonDetails.Name+"...")

	if attempt < chance {
		fmt.Println(pokemonDetails.Name, "was caught!")
		cfg.pokedex[pokemonDetails.Name] = pokemonDetails
	} else {
		fmt.Println(pokemonDetails.Name, "escaped!")
	}

	return nil

}
