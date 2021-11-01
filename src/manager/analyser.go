package main

import (
	"fmt"
	"log"
	"strings"
)

var temperatureMinCheckList []int
var temperatureMaxCheckList []int

var humidityMinCheckList []int
var humidityMaxCheckList []int

var soilMoistureMinCheckList []int
var soilMoistureMaxCheckList []int

var soilTemperatureMinCheckList []int
var soilTemperatureMaxCheckList []int

// HandleAnalyse ...
func HandleAnalyse(min int, max int, current int, sensor string) {
	switch {
	case current > max:
		setMax(0, sensor)
		setMin(1, sensor)
	case current < min:
		setMin(0, sensor)
		setMax(1, sensor)
	case current > min, current < max:
		setMin(1, sensor)
		setMax(1, sensor)
	}
	log.Printf("Analyse %s results", strings.ToLower(sensor))
	log.Printf("The current %s results are %d %% above the minimum treshold.", strings.ToLower(sensor), AnalyseMin(sensor))
	log.Printf("The current %s results are %d %% below the maximum treshold.", strings.ToLower(sensor), AnalyseMax(sensor))
}

// SetMin ...
func setMin(v int, s string) {
	switch s {
	case "Temperature":
		temperatureMinCheckList = append(temperatureMinCheckList, v)
	case "Humidity":
		humidityMinCheckList = append(humidityMinCheckList, v)
	case "SoilMoisture":
		soilMoistureMinCheckList = append(soilMoistureMinCheckList, v)
	case "SoilTemperature":
		soilTemperatureMinCheckList = append(soilTemperatureMinCheckList, v)
	}
}

// SetMax ...
func setMax(v int, s string) {
	switch s {
	case "Temperature":
		temperatureMaxCheckList = append(temperatureMaxCheckList, v)
	case "Humidity":
		humidityMaxCheckList = append(humidityMaxCheckList, v)
	case "SoilMoisture":
		soilMoistureMaxCheckList = append(soilMoistureMaxCheckList, v)
	case "SoilTemperature":
		soilTemperatureMaxCheckList = append(soilTemperatureMaxCheckList, v)
	}
}

// AnalyseMin ...
func AnalyseMin(s string) int {
	var result int
	switch s {
	case "Temperature":
		result = analyse(temperatureMinCheckList)
	case "Humidity":
		result = analyse(humidityMinCheckList)
	case "SoilMoisture":
		result = analyse(soilMoistureMinCheckList)
	case "SoilTemperature":
		result = analyse(soilTemperatureMinCheckList)
	}
	return result
}

// AnalyseMax ...
func AnalyseMax(s string) int {
	var result int
	switch s {
	case "Temperature":
		result = analyse(temperatureMaxCheckList)
	case "Humidity":
		result = analyse(humidityMaxCheckList)
	case "SoilMoisture":
		result = analyse(soilMoistureMaxCheckList)
	case "SoilTemperature":
		result = analyse(soilTemperatureMaxCheckList)
	}
	return result
}

// CleanAnalysis ...
func CleanAnalysis() {
	temperatureMinCheckList = nil
	temperatureMaxCheckList = nil

	humidityMinCheckList = nil
	humidityMaxCheckList = nil

	soilMoistureMinCheckList = nil
	soilMoistureMaxCheckList = nil

	soilTemperatureMinCheckList = nil
	soilTemperatureMaxCheckList = nil
}

func analyse(list []int) int {
	capacity := len(list)
	current := sum(list)
	total := current * 100 / capacity
	return total
}

func sum(list []int) int {
	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum
}

// PrintAnalysisLists ...
func PrintAnalysisLists() {
	fmt.Printf("temperatureMinCheckList %v\n", temperatureMinCheckList)
	fmt.Printf("temperatureMaxCheckList %v\n", temperatureMaxCheckList)

	fmt.Printf("humidityMinCheckList %v\n", humidityMinCheckList)
	fmt.Printf("humidityMaxCheckList %v\n", humidityMaxCheckList)

	fmt.Printf("soilMoistureMinCheckList %v\n", soilMoistureMinCheckList)
	fmt.Printf("soilMoistureMaxCheckList %v\n", soilMoistureMaxCheckList)

	fmt.Printf("soilTemperatureMinCheckList %v\n", soilTemperatureMinCheckList)
	fmt.Printf("soilTemperatureMaxCheckList %v\n", soilTemperatureMaxCheckList)
}
