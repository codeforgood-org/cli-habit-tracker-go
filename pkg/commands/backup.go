// Package commands implements CLI command handlers.
package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Backup creates a backup of the habits data file.
func Backup(store storage.Storage, backupPath string) error {
	// Load habits to ensure file is valid
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	if len(habits) == 0 {
		return fmt.Errorf("no habits to backup")
	}

	// Generate backup filename if not specified
	if backupPath == "" {
		timestamp := time.Now().Format("20060102-150405")
		backupPath = fmt.Sprintf("habits-backup-%s.json", timestamp)
	}

	// Ensure backup directory exists
	dir := filepath.Dir(backupPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// Get source file path
	sourcePath := store.GetPath()

	// Read source file
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Write backup file
	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write backup file: %w", err)
	}

	fmt.Printf("✓ Backup created: %s (%d habit(s))\n", backupPath, len(habits))
	return nil
}

// Restore restores habits from a backup file.
func Restore(store storage.Storage, backupPath string) error {
	// Check if backup file exists
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup file not found: %s", backupPath)
	}

	// Create temporary storage to validate the backup
	tempStorage := storage.NewJSONStorage(backupPath)
	habits, err := tempStorage.Load()
	if err != nil {
		return fmt.Errorf("invalid backup file: %w", err)
	}

	// Validate habits
	for i, habit := range habits {
		if err := habit.Validate(); err != nil {
			return fmt.Errorf("invalid habit at index %d in backup: %w", i, err)
		}
	}

	// Create backup of current data before restoring
	if store.Exists() {
		timestamp := time.Now().Format("20060102-150405")
		autoBackupPath := fmt.Sprintf("habits-auto-backup-%s.json", timestamp)
		currentData, _ := os.ReadFile(store.GetPath())
		if len(currentData) > 0 {
			os.WriteFile(autoBackupPath, currentData, 0644)
			fmt.Printf("Current data backed up to: %s\n", autoBackupPath)
		}
	}

	// Restore from backup
	if err := store.Save(habits); err != nil {
		return fmt.Errorf("failed to restore from backup: %w", err)
	}

	fmt.Printf("✓ Restored %d habit(s) from %s\n", len(habits), backupPath)
	return nil
}
