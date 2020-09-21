package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	configObj *Config
)

// Init config file
// configFilePath: config file path
// Return values
// error if exists
func Init(configFilePath string) error {
	configConent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(configConent), &configObj)
	if err != nil {
		return err
	}

	fmt.Printf("ConfigObj: %v\n", configObj)
	return nil
}

func getConfig() *Config {
	return configObj
}
