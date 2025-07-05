package main

import (
	"fmt"
)

func commandCatch(cfg *config, pokemonName string) error {
	pokemonRes, err := cfg.pokeapiClient.PokemonRetrieve(pokemonName)
	if err != nil {
		return err
	}

	name := pokemonRes.Name
	baseExp := pokemonRes.BaseExp

	fmt.Printf("Pokemon: %s - %v\n", name, baseExp)

	return nil
}
