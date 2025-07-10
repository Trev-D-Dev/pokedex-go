package pokedex

import (
	"sync"
)

type Pokemon struct {
	name string
}

type Pokedex struct {
	mu      sync.Mutex
	pokemon map[string]Pokemon
}

func (p *Pokedex) Add(key string, val []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()

}
