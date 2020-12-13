package main

var temperatureMinCheckList []string
var temperatureMaxCheckList []string

SetMinTemperature(v int){
	append(temperatureMinCheckList, v)
}

SetMaxTemperature(v int){
	append(temperatureMaxCheckList, v)
}

AnalyseMinTemperature() int {
	capacity := len(temperatureMinCheckList)
	current := sum(temperatureMinCheckList)
	total := current * 100 / capacity
	return total
}

AnalyseMinTemperature() int {
	capacity := len(temperatureMinCheckList)
	current := sum(temperatureMinCheckList)
	total := current * 100 / capacity
	return total
}