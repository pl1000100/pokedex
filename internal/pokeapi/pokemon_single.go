package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *ApiClient) SinglePokemon(name string) (Pokemon, error) {
	url := ApiAddress + "/pokemon/" + name
	if data, ok := client.cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
