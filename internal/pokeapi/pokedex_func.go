package pokeapi

import (
	"github.com/Trev-D-Dev/pokedex-go/internal/pokedex"
	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

func (c *Client) AddToPokedex(pokemonInfo types.RespPokemonInfo) error {
	c.pokedex.Add(pokemonInfo)
	return nil
}

func (c *Client) ReturnPokedexEntries() map[string]pokedex.Pokemon {
	return c.pokedex.ReturnEntries()
}
