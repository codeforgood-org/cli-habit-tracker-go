// Package config manages application configuration.
package config

import (
	"os"
	"path/filepath"
)

// Config holds application configuration.
type Config struct {
	DataFilePath string
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		DataFilePath: getDefaultDataFilePath(),
	}
}

// getDefaultDataFilePath returns the default path for the habits data file.
// It tries to use the user's home directory, falling back to current directory.
func getDefaultDataFilePath() string {
	// Try to get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fall back to current directory
		return "habits.json"
	}

	// Create .habit-tracker directory in home
	configDir := filepath.Join(homeDir, ".habit-tracker")
	return filepath.Join(configDir, "habits.json")
}

// New creates a new configuration with optional overrides.
func New(dataFilePath string) *Config {
	cfg := Default()
	if dataFilePath != "" {
		cfg.DataFilePath = dataFilePath
	}
	return cfg
}

// FromEnv creates configuration from environment variables.
func FromEnv() *Config {
	cfg := Default()

	// Override with environment variable if set
	if path := os.Getenv("HABIT_DATA_FILE"); path != "" {
		cfg.DataFilePath = path
	}

	return cfg
}
