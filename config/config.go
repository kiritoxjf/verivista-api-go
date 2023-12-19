package config

import (
	"encoding/json"
	"fmt"
	"os"
	"verivista-api-go/interfaces"
)

// GetConfig 获取配置信息
func GetConfig() (interfaces.Config, error) {
	configPath := os.Getenv("POEM_CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.json"
	}
	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		return interfaces.Config{}, fmt.Errorf("[Config File Error or Not Exist]: %v", err)
	}

	var config interfaces.Config

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return interfaces.Config{}, fmt.Errorf("[Config Analyze Error]: %v", err)
	}

	return config, nil
}
