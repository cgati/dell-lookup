package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ApiKey string
}

func getAPIKey() string {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal(err)
	}

	configuration := Configuration{}
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration.ApiKey
}
