package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var temperatureEndpoint string
var humidityEndpoint string
var soilMoistureEndpoint string
var soilTemperatureEndpoint string

// SensorData ...
type SensorData struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
	Time  string `json:"Time"`
}

func main() {
	config, err := Config()
	if err != nil {
		log.Printf("Cannot load config: %s", err.Error())
	}
	temperatureEndpoint = config.Apis.Temperature.Endpoint
	humidityEndpoint = config.Apis.Humidity.Endpoint
	soilMoistureEndpoint = config.Apis.SoilMoisture.Endpoint
	soilTemperatureEndpoint = config.Apis.SoilTemperature.Endpoint

	temperature, _ := getSensordata(temperatureEndpoint)
	humidity, _ := getSensordata(humidityEndpoint)
	soilMoisture, _ := getSensordata(soilMoistureEndpoint)
	soilTemperature, _ := getSensordata(soilTemperatureEndpoint)
	log.Println(temperature)
	log.Println(humidity)
	log.Println(soilMoisture)
	log.Println(soilTemperature)

}

func getSensordata(endpoint string) (SensorData, error) {
	response, err := http.Get(endpoint)
	responseData, err := ioutil.ReadAll(response.Body)
	var responseObject SensorData
	json.Unmarshal(responseData, &responseObject)
	return responseObject, err
}
