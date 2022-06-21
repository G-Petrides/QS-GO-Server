package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ApplicationId     string
	ApplicationSecret string
	Token             string
	DbAddress         string
	DbPort            string
}

var File = "config.json"

func GetConfig() (Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(File)

	fmt.Println(string(file))

	if err != nil {
		return config, err
	}

	_ = json.Unmarshal(file, &config)

	return config, nil
}
