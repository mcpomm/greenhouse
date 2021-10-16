package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/yryz/ds18b20"
)

// Temperature ...
type Temperature struct {
	ID    string
	Name  string
	Value string
	Unit  string
	Time  string
}

var sensor string

func temperature(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("sensor IDs: %v\n", sensor)

	now := time.Now()
	sec := now.Unix()

	t, err := ds18b20.Temperature(sensor)
	if err == nil {
		temperature := Temperature{sensor, "Temperature", strconv.FormatFloat(t, 'f', 2, 64), "°C", strconv.FormatInt(sec, 10)}

		js, jserr := json.Marshal(temperature)
		if jserr != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func main() {
	// read sensor id from sys/bus/w1/devices/
	sensor = "28-0215c2c3c3ff"

	http.HandleFunc("/", temperature)
	http.ListenAndServe(":5000", nil)

}