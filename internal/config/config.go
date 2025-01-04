package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	configFile = "config.yaml"
)

func Get() *Config {
	return configInstance
}

func LoadConfig() error {
	config, err := readConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	configInstance = config

	// Debug:  loaded configuration
	fmt.Printf("Loaded configuration: %+v\n", *configInstance)

	return nil
}

func readConfig() (*Config, error) {
	var config *Config

	file, err := os.ReadFile(configFile)
	if err != nil {
		return config, fmt.Errorf("error reading file: %v", err)
	}

	// Log file debugging
	//fmt.Printf("Config file content (raw):\n%s\n", string(file))

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling file: %v", err)
	}

	// Log unmarshaled Config struct for debugging
	fmt.Printf("Config struct (parsed):\n%+v\n", config)

	return config, nil
}

func GetDatabaseValues() *Database {
	if configInstance == nil {
		fmt.Println("Config not loaded. Call LoadConfig first.")
		return nil
	}
	return &configInstance.Database
}
