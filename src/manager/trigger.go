package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// TriggerFans ...
func TriggerFans(config Configuration) {
	_trigger("fan", config.Apis.Fan.Endpoint, config.Apis.Fan.Duration)
}

// TriggerAirHeating ...
func TriggerAirHeating(config Configuration) {
	_trigger("air heating", config.Apis.Heating.Endpoint, config.Apis.Heating.Duration)
}

// TriggerWaterTankFill ...
func TriggerWaterTankFill(config Configuration) {
	_trigger("tank filling", config.Apis.WaterPump02.Endpoint, config.Apis.WaterPump02.Duration)
}

// TriggerSoilWatering ...
func TriggerSoilWatering(config Configuration) {
	_trigger("soil watering", config.Apis.WaterPump01.Endpoint, config.Apis.WaterPump01.Duration)
}

// TriggerSoilHeating ...
func TriggerSoilHeating(config Configuration) {
	_trigger("soil heating", config.Apis.HeatingPad.Endpoint, config.Apis.HeatingPad.Duration)
}

func _trigger(action string, endpoint string, duration int) {
	log.Printf("Trigger %s.", action)
	values := map[string]int{"duration": duration}
	jsonData, _ := json.Marshal(values)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	log.Println(res)
}
