package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func List_location_area(url string, client *ApiClient) error {
	if url == "" {
		url = ApiAddress + "/location-area"
	}
	if data, ok := client.Cache.Get(url); ok {
		var locationArea locationAreaResponse
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return err
		}
		for _, result := range locationArea.Results {
			fmt.Println(result.Name)
		}
		client.Next = locationArea.Next
		client.Previous = locationArea.Previous
		fmt.Println("****************************************** From cache")
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	client.Cache.Add(url, body)
	var locationArea locationAreaResponse
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	client.Next = locationArea.Next
	client.Previous = locationArea.Previous
	fmt.Println("****************************************** Not from cache")
	return nil
}
