package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *Config) error {
	res, err := http.Get(*config.Previous)
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
	config.Next = locationArea.Next
	config.Previous = locationArea.Previous
	return nil
}
