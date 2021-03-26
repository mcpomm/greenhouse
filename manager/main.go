package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	config, jsonconfig, err := Config()
	if err != nil {
		log.Printf("Cannot load config: %s", err.Error())
	}
	runSensorCheck(config, jsonconfig)
}

func runSensorCheck(config Configuration, j []byte) {
	counter := 0

	for range time.Tick(config.Monitoring.CheckIntervalMinutes * time.Minute) {
		counter++

		temperature, _ := Sensor{Name: "Temperature"}.GetData(j)
		humidity, _ := Sensor{Name: "Humidity"}.GetData(j)
		soilMoisture, _ := Sensor{Name: "SoilMoisture"}.GetData(j)
		soilTemperature, _ := Sensor{Name: "SoilTemperature"}.GetData(j)

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

func handleSensordata(s string, d Sensor, c Configuration) {
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
			TriggerFans(c)
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The temperature must be increased.")
			TriggerAirHeating(c)
		} else {
			log.Println("The temperature is within the treshold limits.")
		}
	case "Humidity":
		if evaluationMax <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The humidity must be reduced.")
			TriggerFans(c)
		} else if evaluationMin <= c.Monitoring.TresholdLimitPercentage {
			log.Println("The humidity must be increased.")
			TriggerWaterTankFill(c)
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
			TriggerSoilHeating(c)
		} else {
			log.Println("The soil temperature is within the treshold limits.")
		}
	}
}
