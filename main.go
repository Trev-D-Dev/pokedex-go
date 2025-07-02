package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()

			cInput := cleanInput(input)

			word1 := cInput[0]

			if val, ok := commands[word1]; ok {
				err := val.callback()

				if err != nil {
					fmt.Printf("error: %v", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
