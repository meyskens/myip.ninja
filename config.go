package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Bind      string   `json:"bind"`
	Hostnames []string `json:"hostnames"`
	TLS       bool     `json:"tls"`
	CORS      []string `json:"cors"`
}

func getConfig() config {
	returnConfig := config{}

	data, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(data)
	err = jsonParser.Decode(&returnConfig)
	if err != nil {
		panic(err)
	}

	return returnConfig
}
