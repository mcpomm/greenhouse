package main

var temperatureMinCheckList []int
var temperatureMaxCheckList []int

SetMinTemperature(v int){
	append(temperatureMinCheckList, v)
}

SetMaxTemperature(v int){
	append(temperatureMaxCheckList, v)
}

AnalyseMinTemperature() int {
	result := analyse(temperatureMinCheckList)
	return result
}

AnalyseMaxTemperature() int {
	result := analyse(temperatureMaxCheckList)
	return result
}

analyse(list []int) int {
	capacity := len(list)
	current := sum(list)
	total := current * 100 / capacity
	return total
}