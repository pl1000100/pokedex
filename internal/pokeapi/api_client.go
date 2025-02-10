package pokeapi

import "github.com/pl1000100/pokedex/internal/pokecache"

type ApiClient struct {
	Cache    *pokecache.Cache
	Next     string
	Previous string
}
