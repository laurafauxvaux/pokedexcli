package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.next == nil && cfg.previous != nil {
		fmt.Println("You're on the last page")
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil && cfg.next != nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	return nil
}
