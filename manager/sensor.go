package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// SensorData ...
type Sensor struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Unit  string `json:"Unit"`
	Time  string `json:"Time"`
}

func (s Sensor) GetEndpoint(jsonConfig []byte) string {
	search := "Apis.Sensors." + string(s.Name) + ".Endpoint"
	endpoint := gjson.Get(string(jsonConfig), search)
	return endpoint.String()
}

func (s Sensor) GetData(jsonConfig []byte) (Sensor, error) {
	endpoint := s.GetEndpoint(jsonConfig)
	response, err := http.Get(endpoint)
	responseData, err := ioutil.ReadAll(response.Body)

	var responseObject Sensor
	json.Unmarshal(responseData, &responseObject)
	return responseObject, err
}

func (s Sensor) GetTresholdValues() {}
