package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	byteValue, _ := io.ReadAll(data)

	var config Config

	json.Unmarshal(byteValue, &config)

	return &config, nil
}
