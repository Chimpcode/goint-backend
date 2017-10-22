package utils

import (
	"encoding/json"
	"io/ioutil"
)

func GetConfiguration(pathOfConf string) (*GointConfig, error) {
	config := new(GointConfig)

	read, err := ioutil.ReadFile(pathOfConf)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(read, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
