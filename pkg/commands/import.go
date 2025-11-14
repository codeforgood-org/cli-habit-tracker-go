// Package commands implements CLI command handlers.
package commands

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Import imports habits from various formats (CSV, JSON).
func Import(store storage.Storage, format, inputPath string, merge bool) error {
	// Validate format
	format = strings.ToLower(strings.TrimSpace(format))
	if format != "csv" && format != "json" {
		return fmt.Errorf("unsupported format '%s'. Supported formats: csv, json", format)
	}

	// Check if file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", inputPath)
	}

	// Import based on format
	var importedHabits models.HabitList
	var err error

	switch format {
	case "csv":
		importedHabits, err = importCSV(inputPath)
	case "json":
		importedHabits, err = importJSON(inputPath)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return err
	}

	// Validate imported habits
	for i, habit := range importedHabits {
		if err := habit.Validate(); err != nil {
			return fmt.Errorf("invalid habit at row %d: %w", i+1, err)
		}
	}

	if len(importedHabits) == 0 {
		return fmt.Errorf("no habits found in file")
	}

	// Handle merge vs replace
	if merge {
		// Load existing habits
		existingHabits, err := store.Load()
		if err != nil {
			return fmt.Errorf("failed to load existing habits: %w", err)
		}

		// Merge: update existing, add new
		merged := 0
		added := 0
		for _, imported := range importedHabits {
			existing, index := existingHabits.Find(imported.Name)
			if existing != nil {
				// Update existing habit
				existingHabits[index] = imported
				merged++
			} else {
				// Add new habit
				existingHabits = append(existingHabits, imported)
				added++
			}
		}

		// Save merged habits
		if err := store.Save(existingHabits); err != nil {
			return fmt.Errorf("failed to save habits: %w", err)
		}

		fmt.Printf("✓ Imported %d habit(s): %d merged, %d added\n", len(importedHabits), merged, added)
	} else {
		// Replace all habits
		if err := store.Save(importedHabits); err != nil {
			return fmt.Errorf("failed to save habits: %w", err)
		}

		fmt.Printf("✓ Imported %d habit(s) (replaced existing data)\n", len(importedHabits))
	}

	return nil
}

func importCSV(inputPath string) (models.HabitList, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty CSV file")
	}

	// Check if first row is header
	startRow := 0
	if len(records) > 0 && (records[0][0] == "Name" || records[0][0] == "name") {
		startRow = 1
	}

	var habits models.HabitList
	for i := startRow; i < len(records); i++ {
		record := records[i]
		if len(record) < 3 {
			return nil, fmt.Errorf("invalid CSV row %d: expected 3 columns, got %d", i+1, len(record))
		}

		name := strings.TrimSpace(record[0])
		lastDone := strings.TrimSpace(record[1])
		if lastDone == "Never" || lastDone == "" {
			lastDone = ""
		}

		streak, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			return nil, fmt.Errorf("invalid streak value at row %d: %w", i+1, err)
		}

		habit := models.Habit{
			Name:     name,
			LastDone: lastDone,
			Streak:   streak,
		}

		habits = append(habits, habit)
	}

	return habits, nil
}

func importJSON(inputPath string) (models.HabitList, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var habits models.HabitList
	if err := json.Unmarshal(data, &habits); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return habits, nil
}
