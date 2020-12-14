package main

var temperatureMinCheckList []int
var temperatureMaxCheckList []int

var humidityMinCheckList []int
var humidityMaxCheckList []int

var soilMoistureMinCheckList []int
var soilMoistureMaxCheckList []int

var soilTemperatureMinCheckList []int
var soilTemperatureMaxCheckList []int

// SetMin ...
func SetMin(v int, s string) {
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
func SetMax(v int, s string) {
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