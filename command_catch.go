package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, pokemonName string) error {

	currentPokedex := cfg.pokeapiClient.ReturnPokedexEntries()

	_, ok := currentPokedex[pokemonName]

	if ok {
		fmt.Printf("%s is already in Pokedex!\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonRes, err := cfg.pokeapiClient.PokemonRetrieve(pokemonName)
	if err != nil {
		return err
	}

	name := pokemonRes.Name
	baseExp := pokemonRes.BaseExp

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	chance := r.Intn(baseExp)

	minVal := int(baseExp - (baseExp / 10))

	if chance >= minVal {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokeapiClient.AddToPokedex(pokemonRes)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
