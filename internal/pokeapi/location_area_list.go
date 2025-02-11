package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *ApiClient) ListLocationArea(url *string) (locationAreas, error) {
	useUrl := ApiAddress + "/location-area"
	if url != nil {
		useUrl = *url
	}
	if data, ok := client.cache.Get(useUrl); ok {
		var locationArea locationAreas
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return locationAreas{}, err
		}
		return locationArea, nil
	}
	res, err := http.Get(useUrl)
	if err != nil {
		return locationAreas{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return locationAreas{}, err
	}
	client.cache.Add(useUrl, body)
	var locationArea locationAreas
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return locationAreas{}, err
	}
	return locationArea, nil
}
