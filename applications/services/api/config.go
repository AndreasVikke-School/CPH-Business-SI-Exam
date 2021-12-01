package main

import (
	"encoding/json"
	"os"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/api/errorhandler"
)

type Configuration struct {
	Postgres struct {
		Service string `json:"service"`
	} `json:"postgres"`
	Redis struct {
		Service string `json:"service"`
	} `json:"redis"`
	Kafka struct {
		Service string `json:"service"`
	} `json:"kafka"`
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
