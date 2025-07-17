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
	pExp := pokemonRes.BaseExp

	baseExp := 40.0
	pExpF := float64(pExp)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	chance := r.Float64() * pExpF

	divFactor := 0.0

	if pExp < 125 {
		divFactor = 75.0
	} else {
		divFactor = 50.0
	}

	minVal := baseExp * (pExpF / divFactor)

	if chance >= minVal {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokeapiClient.AddToPokedex(pokemonRes)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
