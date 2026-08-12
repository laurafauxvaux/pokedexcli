package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
			"exit": {
				name: "exit", 
				description: "Exit the Pokedex",
				callback: commandExit,
			},
			"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
		}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()

	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}