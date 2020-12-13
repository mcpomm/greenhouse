package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	handleSensordata("Temperature", temperature, config)
	handleSensordata("Humidity", humidity, config)
	handleSensordata("SoilMoisture", soilMoisture, config)
	handleSensordata("SoilTemperature", soilTemperature, config)

}

func getSensordata(endpoint string) (SensorData, error) {
	response, err := http.Get(endpoint)
	responseData, err := ioutil.ReadAll(response.Body)
	var responseObject SensorData
	json.Unmarshal(responseData, &responseObject)
	return responseObject, err
}

func handleSensordata(s string, d SensorData, c Configuration) {
	minValue, maxValue := GetTresholdValues(s, &c)
	fmt.Println()
	log.Println("Check", s)
	log.Println("----------------------")
	log.Printf("minimum value: %d %s\n", minValue, d.Unit)
	log.Printf("maximum value: %d %s\n", maxValue, d.Unit)
	log.Printf("current value: %s %s\n", d.Value, d.Unit)

	switch s {
	case "Temperature":
		ti, _ := strconv.ParseFloat(d.Value, 64)
		handleTemperature(minValue, maxValue, int(ti))
	case "Humidity":
		handleHumidity()
	case "SoilMoisture":
		handleSoilMoisture()
	case "SoilTemperature":
		handleSoilTemperature()
	}
}

func handleTemperature(min int, max int, current int) {
	switch {
	case current < min:
		SetMinTemperature(0)
		SetMaxTemperature(1)
	case current > min, current < max:
		SetMinTemperature(1)
		SetMaxTemperature(1)
	case current > max:
		SetMaxTemperature(0)
	}
	log.Println("Analyse temperature results")
	log.Printf("The current temperature results are %d %% above the minimum treshold.", AnalyseMinTemperature())
	log.Printf("The current temperature results are %d %% below the maximum treshold.", AnalyseMaxTemperature())
}

func handleHumidity() {
	fmt.Println("handle humidity")
}

func handleSoilMoisture() {
	fmt.Println("handle soilMoisture")
}

func handleSoilTemperature() {
	fmt.Println("handle soilTemperature")
}
