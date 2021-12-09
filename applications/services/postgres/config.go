package main

import (
	"encoding/json"
	"os"

	eh "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/postgres/errorhandler"
)

type Configuration struct {
	Postgres struct {
		Broker   string `json:"broker"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"postgres"`
}

func getConfig(env string) Configuration {
	file, _ := os.Open("configs/" + env + "_conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	eh.PanicOnError(err, "Failed to Decode Env File")
	return configuration
}
