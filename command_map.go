package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    uint
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	var locationArea locationAreaResponse
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}
