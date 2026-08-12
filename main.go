package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			finalInput := cleanInput(userInput)
			
			if len(finalInput) > 0 {
				command := finalInput[0]
				if cmd, ok := commands[command]; ok {
					if err := cmd.callback(); err != nil {
						fmt.Printf("%v\n", err)
					} 	
				} else {
					fmt.Print("Unknown command\n")
				}
		}
	}
}
}