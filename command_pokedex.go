package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {

	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("You have yet to catch a Pokemon...")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Println("-" + name)
	}

	return nil
}
