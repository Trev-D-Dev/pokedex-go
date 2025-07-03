package main

import (
	"time"

	"github.com/Trev-D-Dev/pokedex-go/internal/pokeapi"
)

func main() {
	pClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pClient,
	}

	startRepl(cfg)
}
