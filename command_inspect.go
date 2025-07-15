package main

import (
	"fmt"

	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

func commandInspect(cfg *config, pokemonName string) error {

	currentPokedex := cfg.pokeapiClient.ReturnPokedexEntries()

	var pokemon types.Pokemon
	pokemon, ok := currentPokedex[pokemonName]

	if !ok {
		fmt.Printf("%s is not in Pokedex!\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)

	fmt.Printf("Height: %v\n", pokemon.Height)

	fmt.Println("Stats:")

	for stat, val := range pokemon.Stats {
		fmt.Printf("   -%s: %v\n", stat, val)
	}

	fmt.Println("Types:")

	for _, pType := range pokemon.Types {
		fmt.Printf("   - %s\n", pType)
	}

	return nil
}
