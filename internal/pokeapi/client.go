package pokeapi

import (
	"net/http"
	"time"

	"github.com/mattnickolaus/pokedexREPL/internal/pokecache"
)

// Client
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// Create New Client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
