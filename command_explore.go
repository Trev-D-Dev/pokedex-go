package main

/*
import (
	"errors"
)*/

import "fmt"

func commandExplore(cfg *config, areaName string) error {
	fmt.Printf("Exploring %s...\n", areaName)

	pokemonRes, err := cfg.pokeapiClient.ListPokemon(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range pokemonRes.PokeEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
