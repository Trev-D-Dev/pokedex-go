package pokeapi

import (
	"net/http"
	"time"

	"github.com/Trev-D-Dev/pokedex-go/internal/pokecache"
)

// Client struct
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient function
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(timeout),
	}
}
