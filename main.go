package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()
		finalInput := cleanInput(userInput)

		if len(finalInput) > 0 {
			command := finalInput[0]
			if cmd, ok := cfg.commands[command]; ok {
				if err := cmd.callback(cfg); err != nil {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Print("Unknown command\n")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("%s\n", err)
	}
}

func main() {
	cfg := &config{
		commands: getCommands(),
	}
	repl(cfg)
}
