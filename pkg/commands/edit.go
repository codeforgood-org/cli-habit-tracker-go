// Package commands implements CLI command handlers.
package commands

import (
	"fmt"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Edit renames a habit.
func Edit(store storage.Storage, oldName, newName string) error {
	// Validate input
	oldName = strings.TrimSpace(oldName)
	newName = strings.TrimSpace(newName)

	if oldName == "" {
		return fmt.Errorf("current habit name cannot be empty")
	}
	if newName == "" {
		return fmt.Errorf("new habit name cannot be empty")
	}

	// Load existing habits
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	// Find the habit to edit
	habit, index := habits.Find(oldName)
	if habit == nil {
		return fmt.Errorf("habit '%s' not found", oldName)
	}

	// Check if new name already exists
	if existing, _ := habits.Find(newName); existing != nil && !strings.EqualFold(oldName, newName) {
		return fmt.Errorf("habit '%s' already exists", newName)
	}

	// Update the habit name
	habit.Name = newName
	habits[index] = *habit

	// Save updated habits
	if err := store.Save(habits); err != nil {
		return fmt.Errorf("failed to save habits: %w", err)
	}

	fmt.Printf("✓ Renamed habit '%s' to '%s'\n", oldName, newName)
	return nil
}
