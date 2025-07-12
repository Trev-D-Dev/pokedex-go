package pokedex

import (
	"sync"

	"github.com/Trev-D-Dev/pokedex-go/internal/types"
)

type Pokemon struct {
	name    string
	baseEXP int
	height  int
	weight  int
	stats   map[string]int
	types   []string
}

type Pokedex struct {
	mu      sync.Mutex
	pokemon map[string]Pokemon
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
	types := make([]string, lenTypes)
	for i, pType := range pokemon.Types {
		tName := pType.Type.Name
		types[i] = tName
	}

	newPokemon := Pokemon{
		name:    key,
		baseEXP: pokemon.BaseExp,
		height:  pokemon.Height,
		weight:  pokemon.Weight,
		stats:   stats,
		types:   types,
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.pokemon[key] = newPokemon
}

func NewPokedex() *Pokedex {
	newPokedex := Pokedex{}
	newPokedex.pokemon = make(map[string]Pokemon)

	return &newPokedex
}

func (p *Pokedex) Get(key string) (Pokemon, bool) {
	entry, ok := p.pokemon[key]
	if !ok {
		empty := Pokemon{}
		return empty, false
	} else {
		return entry, true
	}
}

func (p *Pokedex) ReturnEntries() map[string]Pokemon {
	return p.pokemon
}
