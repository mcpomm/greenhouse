package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var temperatureEndpoint string
var humidityEndpoint string
var soilMoistureEndpoint string
var soilTemperatureEndpoint string

var c chan int

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
	runSensorCheck(config)

}

func runSensorCheck(config Configuration) {
	counter := 0

	temperatureEndpoint = config.Apis.Temperature.Endpoint
	humidityEndpoint = config.Apis.Humidity.Endpoint
	soilMoistureEndpoint = config.Apis.SoilMoisture.Endpoint
	soilTemperatureEndpoint = config.Apis.SoilTemperature.Endpoint

	for range time.Tick(10 * time.Second) {
		counter++

		fmt.Println(counter)

		temperature, _ := getSensordata(temperatureEndpoint)
		humidity, _ := getSensordata(humidityEndpoint)
		soilMoisture, _ := getSensordata(soilMoistureEndpoint)
		soilTemperature, _ := getSensordata(soilTemperatureEndpoint)
		handleSensordata("Temperature", temperature, config)
		handleSensordata("Humidity", humidity, config)
		handleSensordata("SoilMoisture", soilMoisture, config)
		handleSensordata("SoilTemperature", soilTemperature, config)

		PrintAnalysisLists()

		if counter == config.Monitoring.CheckIntervalCountPerEvaluation {
			counter = 0
			CleanAnalysis()
		}
	}
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

	vi, _ := strconv.ParseFloat(d.Value, 64)
	handleAnalyse(minValue, maxValue, int(vi), s)
}

func handleAnalyse(min int, max int, current int, sensor string) {
	switch {
	case current < min:
		SetMin(0, sensor)
		SetMax(1, sensor)
	case current > min, current < max:
		SetMin(1, sensor)
		SetMax(1, sensor)
	case current > max:
		SetMax(0, sensor)
	}
	log.Printf("Analyse %s results", strings.ToLower(sensor))
	log.Printf("The current %s results are %d %% above the minimum treshold.", strings.ToLower(sensor), AnalyseMin(sensor))
	log.Printf("The current %s results are %d %% below the maximum treshold.", strings.ToLower(sensor), AnalyseMax(sensor))
}
