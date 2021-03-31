package server

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Addr         string `json:"addr"`
	Database     string `json:"database"`
	DatabasePath string `json:"databasePath"`
}

var mainConfig *Config

func init() {
	mainConfig = &Config{
		Addr: ":8080",
	}
}

func ReadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, mainConfig); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return mainConfig
}
