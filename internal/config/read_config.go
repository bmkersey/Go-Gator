package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	var cfg Config

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".gatorconfig.json")

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to decode config file: %w", err)
	}

	return cfg, nil
}
