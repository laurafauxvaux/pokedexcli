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

var commands = map[string]cliCommand {
	"exit": {
		name: "exit", 
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}