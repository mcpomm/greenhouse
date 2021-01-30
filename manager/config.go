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
		WaterPump struct {
			Endpoint     string `json:"Endpoint"`
			PumpDuration int    `json:"PumpDuration"`
		} `json:"WaterPump"`
	} `json:"Apis"`
}

// Config ...
func Config() (Configuration, error) {
	byteValue, err := readConfigFile()
	var conf Configuration
	json.Unmarshal(byteValue, &conf)
	return conf, err
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
	min := reflect.Indirect(r).FieldByName("Apis").FieldByName(field).FieldByName("TresholdMin")
	max := reflect.Indirect(r).FieldByName("Apis").FieldByName(field).FieldByName("TresholdMax")
	return int(min.Int()), int(max.Int())
}
