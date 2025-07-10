package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, pokemonName string) error {

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
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
