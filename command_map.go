package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	body, ok := config.cache.Get(*config.next)
	if !ok {
		res, err := http.Get(*config.next)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}
		config.cache.Add(*config.next, body)
		fmt.Println("****************************************** Not from cache")
	} else {
		fmt.Println("****************************************** From cache")
	}
	var locationArea locationAreaResponse
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	config.next = locationArea.Next
	config.previous = locationArea.Previous
	return nil
}

func commandMapb(config *Config) error {
	body, ok := config.cache.Get(*config.previous)
	if !ok {
		res, err := http.Get(*config.previous)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}
		config.cache.Add(*config.previous, body)
		fmt.Println("****************************************** Not from cache")
	} else {
		fmt.Println("****************************************** From cache")
	}
	var locationArea locationAreaResponse
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	config.next = locationArea.Next
	config.previous = locationArea.Previous
	return nil
}
