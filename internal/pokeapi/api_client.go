package pokeapi

import (
	"time"

	"github.com/pl1000100/pokedex/internal/pokecache"
)

type ApiClient struct {
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) ApiClient {
	return ApiClient{
		cache: pokecache.NewCache(cacheInterval),
	}
}
