package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokemonName string) error {

	currentPokedex := cfg.pokeapiClient.ReturnPokedexEntries()

	fmt.Println("Your Pokedex:")

	for _, pokemon := range currentPokedex {
		fmt.Printf("   - %s\t", pokemon.Name)
	}

	return nil
}
