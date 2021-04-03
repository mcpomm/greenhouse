package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Sensor ...
type Sensor struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
	Time  string `json:"Time"`
}

func (s Sensor) GetConfig(c Configuration) reflect.Value {
	r := reflect.ValueOf(c)
	config := reflect.Indirect(r).FieldByName("Apis").FieldByName("Sensors").FieldByName(s.Name)
	return config
}

func (s Sensor) GetData(c Configuration) (Sensor, error) {
	endpoint := s.GetEndpoint(c)
	response, err := http.Get(endpoint)
	responseData, err := ioutil.ReadAll(response.Body)

	var responseObject Sensor
	json.Unmarshal(responseData, &responseObject)
	return responseObject, err
}

func (s Sensor) GetEndpoint(c Configuration) string {
	config := s.GetConfig(c)
	endpoint := config.FieldByName("Endpoint")
	return endpoint.String()
}

func (s Sensor) GetThresholdValues(c Configuration) (int, int) {
	config := s.GetConfig(c)
	min := config.FieldByName("TresholdMin")
	max := config.FieldByName("TresholdMax")
	return int(min.Int()), int(max.Int())
}
