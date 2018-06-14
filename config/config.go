package config

import (
	"encoding/json"
	"os"

	"github.com/qwertypomy/printers/models"
)

func GetConfig() (models.Config, error) {
	config := models.Config{}
	file, err := os.Open("./config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
