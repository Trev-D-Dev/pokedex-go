package pokedex

import (
	"sync"

	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

type Pokedex struct {
	mu      sync.Mutex
	pokemon map[string]types.Pokemon
}

func (p *Pokedex) Add(pokemon types.RespPokemonInfo) {
	key := pokemon.Name

	stats := make(map[string]int)
	for _, stat := range pokemon.Stats {
		base := stat.BaseStat
		name := stat.Stat.Name
		stats[name] = base
	}

	lenTypes := len(pokemon.Types)
	pokeTypes := make([]string, lenTypes)
	for i, pType := range pokemon.Types {
		tName := pType.Type.Name
		pokeTypes[i] = tName
	}

	newPokemon := types.Pokemon{
		Name:    key,
		BaseEXP: pokemon.BaseExp,
		Height:  pokemon.Height,
		Weight:  pokemon.Weight,
		Stats:   stats,
		Types:   pokeTypes,
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.pokemon[key] = newPokemon
}

func NewPokedex() *Pokedex {
	newPokedex := Pokedex{}
	newPokedex.pokemon = make(map[string]types.Pokemon)

	return &newPokedex
}

func (p *Pokedex) Get(key string) (types.Pokemon, bool) {
	entry, ok := p.pokemon[key]
	if !ok {
		empty := types.Pokemon{}
		return empty, false
	} else {
		return entry, true
	}
}

func (p *Pokedex) ReturnEntries() map[string]types.Pokemon {
	return p.pokemon
}
