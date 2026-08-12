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
			cleanedInput := cleanInput(userInput)
			if len(cleanedInput) > 0 {
			firstWord := cleanedInput[0]
			}
		}
	}
}