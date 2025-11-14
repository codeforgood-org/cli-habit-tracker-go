// Package commands implements CLI command handlers.
package commands

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Export exports habits to various formats (CSV, JSON).
func Export(store storage.Storage, format, outputPath string) error {
	// Validate format
	format = strings.ToLower(strings.TrimSpace(format))
	if format != "csv" && format != "json" {
		return fmt.Errorf("unsupported format '%s'. Supported formats: csv, json", format)
	}

	// Load habits
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	if len(habits) == 0 {
		return fmt.Errorf("no habits to export")
	}

	// Create output directory if needed
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Export based on format
	switch format {
	case "csv":
		return exportCSV(habits, outputPath)
	case "json":
		return exportJSON(habits, outputPath)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func exportCSV(habits models.HabitList, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Name", "Last Done", "Streak"}); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write habits
	for _, habit := range habits {
		lastDone := habit.LastDone
		if lastDone == "" {
			lastDone = "Never"
		}
		row := []string{
			habit.Name,
			lastDone,
			strconv.Itoa(habit.Streak),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	fmt.Printf("✓ Exported %d habit(s) to %s\n", len(habits), outputPath)
	return nil
}

func exportJSON(habits models.HabitList, outputPath string) error {
	data, err := json.MarshalIndent(habits, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("✓ Exported %d habit(s) to %s\n", len(habits), outputPath)
	return nil
}
