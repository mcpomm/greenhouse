package main

var temperatureMinCheckList []int
var temperatureMaxCheckList []int

// SetMinTemperature ...
func SetMinTemperature(v int) {
	temperatureMinCheckList = append(temperatureMinCheckList, v)
	// fmt.Println(temperatureMinCheckList)
}

// SetMaxTemperature ...
func SetMaxTemperature(v int) {
	temperatureMaxCheckList = append(temperatureMaxCheckList, v)
	// fmt.Println(temperatureMaxCheckList)
}

// AnalyseMinTemperature ...
func AnalyseMinTemperature() int {
	result := analyse(temperatureMinCheckList)
	return result
}

// AnalyseMaxTemperature ...
func AnalyseMaxTemperature() int {
	result := analyse(temperatureMaxCheckList)
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
