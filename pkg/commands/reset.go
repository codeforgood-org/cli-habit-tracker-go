// Package commands implements CLI command handlers.
package commands

import (
	"fmt"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Reset resets a habit's streak to zero.
func Reset(store storage.Storage, habitName string) error {
	// Validate input
	habitName = strings.TrimSpace(habitName)
	if habitName == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	// Load existing habits
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	// Find the habit
	habit, index := habits.Find(habitName)
	if habit == nil {
		return fmt.Errorf("habit '%s' not found", habitName)
	}

	// Reset the streak
	oldStreak := habit.Streak
	habit.Streak = 0
	habit.LastDone = ""
	habits[index] = *habit

	// Save updated habits
	if err := store.Save(habits); err != nil {
		return fmt.Errorf("failed to save habits: %w", err)
	}

	fmt.Printf("✓ Habit '%s' has been reset (previous streak: %d day(s)).\n", habitName, oldStreak)
	return nil
}
