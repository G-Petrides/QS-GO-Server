package core

import (
	"encoding/json"
	"io/ioutil"
)
//LinnConfig : struct for parsing Linnworks API data
//from the config.json file in the root folder.
type LinnConfig struct {
    ApplicationID     string    `json:"applicationId"`
	ApplicationSecret string
	Token             string
}

//DbConfig : struct for parsing Database connection settings
//from the config.json file in the root folder.
type DbConfig struct {
	DbAddress string
	DbPort    string
}

//Config : struct for unmarshaling config.json
type Config struct {
	Linn LinnConfig
	Db   DbConfig
}

//File is a global setting with the path to config.json
//Allows setting of a relative path for unit testing.
var File = "config.json"

//GetConfig reads config.json, unmarshals and returns the config data.
func GetConfig() (Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(File)

	if err != nil {
		return config, err
	}

	_ = json.Unmarshal(file, &config)

	return config, nil
}
