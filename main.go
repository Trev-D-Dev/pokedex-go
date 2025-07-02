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

			fmt.Printf("Your command was: %v\n", word1)
		}
	}
}
