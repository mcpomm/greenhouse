package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

const defaultConfig = "test"

// Configuration ...
type Configuration struct {
	Monitoring struct {
		ThingSpeakAPI                   string        `json:"ThingSpeakAPI"`
		CheckIntervalMinutes            time.Duration `json:"CheckIntervalMinutes"`
		CheckIntervalCountPerEvaluation int           `json:"CheckIntervalCountPerEvaluation"`
		TresholdLimitPercentage         int           `json:"TresholdLimitPercentage"`
	}
	Apis struct {
		Sensors struct {
			Temperature struct {
				Endpoint    string `json:"Endpoint"`
				TresholdMin int    `json:"TresholdMin"`
				TresholdMax int    `json:"TresholdMax"`
			} `json:"Temperature"`
			Humidity struct {
				Endpoint    string `json:"Endpoint"`
				TresholdMin int    `json:"TresholdMin"`
				TresholdMax int    `json:"TresholdMax"`
			} `json:"Humidity"`
			SoilMoisture struct {
				Endpoint    string `json:"Endpoint"`
				TresholdMin int    `json:"TresholdMin"`
				TresholdMax int    `json:"TresholdMax"`
			} `json:"SoilMoisture"`
			SoilTemperature struct {
				Endpoint    string `json:"Endpoint"`
				TresholdMin int    `json:"TresholdMin"`
				TresholdMax int    `json:"TresholdMax"`
			} `json:"SoilTemperature"`
		} `json:"Sensors"`
		Trigger struct {
			WaterPump01 struct {
				Endpoint string `json:"Endpoint"`
				Duration int    `json:"Duration"`
			} `json:"WaterPump01"`
			HeatingPad struct {
				Endpoint string `json:"Endpoint"`
				Duration int    `json:"Duration"`
			} `json:"HeatingPad"`
			Fan struct {
				Endpoint string `json:"Endpoint"`
				Duration int    `json:"Duration"`
			} `json:"Fan"`
			Heating struct {
				Endpoint string `json:"Endpoint"`
				Duration int    `json:"Duration"`
			} `json:"Heating"`
			WaterPump02 struct {
				Endpoint string `json:"Endpoint"`
				Duration int    `json:"Duration"`
			} `json:"WaterPump02"`
		} `json:Trigger`
	} `json:"Apis"`
}

// Config ...
func Config() (Configuration, []byte, error) {
	jsonFile, err := readConfigFile()
	var conf Configuration
	json.Unmarshal(jsonFile, &conf)
	return conf, jsonFile, err
}

func readConfigFile() ([]byte, error) {
	filepath := filepath.Join("config", getConfigName())
	file, err := os.Open(filepath)
	log.Printf("Successfully opened configuration file %s", filepath)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	return byteValue, err
}

func getConfigName() string {
	if c, present := os.LookupEnv("CONFIG"); present {
		return fmt.Sprintf("%s.json", c)
	}
	return fmt.Sprintf("%s.json", defaultConfig)
}

// GetTresholdValues ...
func GetTresholdValues(field string, c *Configuration) (int, int) {
	r := reflect.ValueOf(c)
	min := reflect.Indirect(r).FieldByName("Apis").FieldByName("Sensors").FieldByName(field).FieldByName("TresholdMin")
	max := reflect.Indirect(r).FieldByName("Apis").FieldByName("Sensors").FieldByName(field).FieldByName("TresholdMax")
	return int(min.Int()), int(max.Int())
}
