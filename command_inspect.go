package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no Pokemon name was given")
	}

	pokemon := args[0]

	pokemonDetails, ok := cfg.pokedex[pokemon]

	if !ok {
		return fmt.Errorf("You have not caught that Pokemon")
	}

	fmt.Println("Name:", pokemonDetails.Name)
	fmt.Println("Height:", pokemonDetails.Height)
	fmt.Println("Weight:", pokemonDetails.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Println("-"+stat.Stat.Name+":", stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemonDetails.Types {
		fmt.Println("-" + t.Type.Name)
	}

	return nil
}
