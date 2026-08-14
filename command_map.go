package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type resources struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}

func fetch(url string) (resources, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resources{}, fmt.Errorf("failed to get response: %w", err)
	}

	defer resp.Body.Close()

	var locations resources

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&locations); err != nil {
		return resources{}, fmt.Errorf("failed to decode: %w", err)
	}
	return locations, nil
}

func commandMap(cfg *config) error {
	if cfg.next == nil {
		fmt.Println("You're on the last page")
		return nil
	}
	locations, err := fetch(*cfg.next)
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
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	locations, err := fetch(*cfg.previous)
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
