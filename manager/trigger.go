package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// TriggerFans ...
func TriggerFans() {
	log.Println("Trigger fans.")
}

// TriggerAirHeating ...
func TriggerAirHeating() {
	log.Println("Trigger air heating.")
}

// TriggerWaterTankFill ...
func TriggerWaterTankFill() {
	log.Println("Trigger water tank fill.")
}

// TriggerSoilWatering ...
func TriggerSoilWatering(config Configuration) {
	log.Println("Trigger soil watering.")
	waterPumpEndpoint := config.Apis.WaterPump.Endpoint
	values := map[string]string{"pumpDuration": config.Apis.WaterPump.PumpDuration}
	jsonData, _ := json.Marshal(values)

	response, err := http.Post(waterPumpEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	log.Println(res["json"])
}

// TriggerSoilHeating ...
func TriggerSoilHeating() {
	log.Println("Trigger soil heating.")
}
