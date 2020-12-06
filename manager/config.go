package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const defaultConfig = "test"

// Configuration ...
type Configuration struct {
	Apis struct {
		Temperature struct {
			Endpoint string `json:"Endpoint"`
		} `json:"Temperature"`
		Humidity struct {
			Endpoint string `json:"Endpoint"`
		} `json:"Humidity"`
		SoilMoisture struct {
			Endpoint string `json:"Endpoint"`
		} `json:"SoilMoisture"`
		SoilTemperature struct {
			Endpoint string `json:"Endpoint"`
		} `json:"SoilTemperature"`
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
