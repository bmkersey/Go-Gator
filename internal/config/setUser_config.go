package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (c *Config) SetUser(username string) error {
	c.CurrentUser = username

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	configPath := filepath.Join(homeDir, ".gatorconfig.json")

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
