package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// TriggerFans ...
func TriggerFans(config Configuration) {
	_trigger("fan", config.Apis.Trigger.Fan.Endpoint, config.Apis.Trigger.Fan.Duration)
}

// TriggerAirHeating ...
func TriggerAirHeating(config Configuration) {
	_trigger("air heating", config.Apis.Trigger.Heating.Endpoint, config.Apis.Trigger.Heating.Duration)
}

// TriggerWaterTankFill ...
func TriggerWaterTankFill(config Configuration) {
	_trigger("tank filling", config.Apis.Trigger.WaterPump02.Endpoint, config.Apis.Trigger.WaterPump02.Duration)
}

// TriggerSoilWatering ...
func TriggerSoilWatering(config Configuration) {
	_trigger("soil watering", config.Apis.Trigger.WaterPump01.Endpoint, config.Apis.Trigger.WaterPump01.Duration)
}

// TriggerSoilHeating ...
func TriggerSoilHeating(config Configuration) {
	_trigger("soil heating", config.Apis.Trigger.HeatingPad.Endpoint, config.Apis.Trigger.HeatingPad.Duration)
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
