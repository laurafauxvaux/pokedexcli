package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location area given")
	}

	areaName := args[0]
	area, err := cfg.pokeapiClient.AreaDetails(areaName)

	if err != nil {
		return err
	}

	for _, encounter := range area.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
