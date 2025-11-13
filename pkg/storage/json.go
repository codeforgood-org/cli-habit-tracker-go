// Package storage handles data persistence for habits.
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
)

// JSONStorage implements habit storage using JSON files.
type JSONStorage struct {
	filePath string
}

// NewJSONStorage creates a new JSON storage instance.
func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{
		filePath: filePath,
	}
}

// Load reads habits from the JSON file.
func (s *JSONStorage) Load() (models.HabitList, error) {
	// Check if file exists
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		// Return empty list if file doesn't exist
		return models.HabitList{}, nil
	}

	// Read file
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Handle empty file
	if len(data) == 0 {
		return models.HabitList{}, nil
	}

	// Unmarshal JSON
	var habits models.HabitList
	if err := json.Unmarshal(data, &habits); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return habits, nil
}

// Save writes habits to the JSON file.
func (s *JSONStorage) Save(habits models.HabitList) error {
	// Ensure directory exists
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal to JSON with indentation
	data, err := json.MarshalIndent(habits, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Delete removes the storage file.
func (s *JSONStorage) Delete() error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}
	return os.Remove(s.filePath)
}

// Exists checks if the storage file exists.
func (s *JSONStorage) Exists() bool {
	_, err := os.Stat(s.filePath)
	return err == nil
}

// GetPath returns the file path.
func (s *JSONStorage) GetPath() string {
	return s.filePath
}

// Storage defines the interface for habit persistence.
type Storage interface {
	Load() (models.HabitList, error)
	Save(models.HabitList) error
	Delete() error
	Exists() bool
	GetPath() string
}
