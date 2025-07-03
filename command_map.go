package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationRes.Next
	cfg.prevURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationRes.Next
	cfg.prevURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
