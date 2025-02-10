package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *ApiClient) SingleLocationArea(location_name string) (location, error) {
	url := ApiAddress + "/location-area/" + location_name
	if data, exist := client.cache.Get(url); exist {
		var loc location
		if err := json.Unmarshal(data, &loc); err != nil {
			return location{}, err
		}
		return loc, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return location{}, nil
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return location{}, nil
	}
	client.cache.Add(url, body)
	var loc location
	if err := json.Unmarshal(body, &loc); err != nil {
		return location{}, err
	}
	return loc, nil
}
