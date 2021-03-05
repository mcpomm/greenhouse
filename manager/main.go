package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
	InitializeTriggerPins()
	runSensorCheck(config)
}

func runSensorCheck(config Configuration) {
	counter := 0

	// Sensor Endpoints
	temperatureEndpoint = config.Apis.Temperature.Endpoint
	humidityEndpoint = config.Apis.Humidity.Endpoint
	soilMoistureEndpoint = config.Apis.SoilMoisture.Endpoint
	soilTemperatureEndpoint = config.Apis.SoilTemperature.Endpoint

	for range time.Tick(config.Monitoring.CheckIntervalMinutes * time.Minute) {
		counter++

		temperature, _ := getSensordata(temperatureEndpoint)
		humidity, _ := getSensordata(humidityEndpoint)
		soilMoisture, _ := getSensordata(soilMoistureEndpoint)
		soilTemperature, _ := getSensordata(soilTemperatureEndpoint)

		ThingSpeak(
			config,
			Payload{
				APIKey: os.Getenv("THING_SPEAK_API_KEY"),
				Field1: temperature.Value,
				Field2: humidity.Value,
				Field3: soilMoisture.Value,
				Field4: soilTemperature.Value})

		handleSensordata("Temperature", temperature, config)
		handleSensordata("Humidity", humidity, config)
		handleSensordata("SoilMoisture", soilMoisture, config)
		handleSensordata("SoilTemperature", soilTemperature, config)

		PrintAnalysisLists()

		if counter == config.Monitoring.CheckIntervalCountPerEvaluation {
			counter = 0
			evaluateAnalysis("Temperature", config)
			evaluateAnalysis("Humidity", config)
			evaluateAnalysis("SoilMoisture", config)
			evaluateAnalysis("SoilTemperature", config)
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
	HandleAnalyse(minValue, maxValue, int(vi), s)
}

func evaluateAnalysis(s string, c Configuration) {
	evaluationMin := AnalyseMin(s)
	evaluationMax := AnalyseMax(s)
	switch s {
	case "Temperature":
		if evaluationMax <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The temperature must be reduced.")
			TriggerFans()
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The temperature must be increased.")
			TriggerAirHeating()

		} else {
			log.Println("The temperature is within the treshold limits.")
		}
	case "Humidity":
		if evaluationMax <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The humidity must be reduced.")
			TriggerFans()
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The humidity must be increased.")
			TriggerWaterTankFill()
		} else {
			log.Println("The humidity is within the treshold limits.")
		}
	case "SoilMoisture":
		if evaluationMax <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The soil moisture must be reduced.")
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The soil moisture must be increased.")
			TriggerSoilWatering(c)
		} else {
			log.Println("The soil moisture is within the treshold limits.")
		}
	case "SoilTemperature":
		if evaluationMax <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The soil temperature must be reduced.")
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The soil temperature must be increased.")
			TriggerSoilHeating()
		} else {
			log.Println("The soil temperature is within the treshold limits.")
		}
	}
}
