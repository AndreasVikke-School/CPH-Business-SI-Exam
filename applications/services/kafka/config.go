package main

import (
	"encoding/json"
	"os"

	eh "github.com/AndreasVikke-School/CPH-Bussiness-SI-Exam/applications/services/kafka/errorhandler"
)

type Configuration struct {
	Kafka struct {
		Broker string `json:"broker"`
	} `json:"kafka"`
	Redis struct {
		Broker string `json:"broker"`
	} `json:"redis"`
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
