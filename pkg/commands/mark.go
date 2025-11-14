// Package commands implements CLI command handlers.
package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Mark marks a habit as completed for today.
func Mark(store storage.Storage, habitName string) error {
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

	today := time.Now()

	// Check if habit exists
	habit, index := habits.Find(habitName)
	if habit != nil {
		// Existing habit - update streak
		err := habit.UpdateStreak(today)
		if err != nil {
			// Already marked today
			fmt.Printf("✓ '%s' is already marked for today!\n", habitName)
			return nil
		}

		// Update the habit in the list
		habits[index] = *habit
		fmt.Printf("✓ Marked '%s' as done today! Current streak: %d day(s)\n", habitName, habit.Streak)
	} else {
		// New habit - create and add
		newHabit := models.Habit{
			Name:     habitName,
			LastDone: today.Format("2006-01-02"),
			Streak:   1,
		}

		if err := newHabit.Validate(); err != nil {
			return fmt.Errorf("invalid habit: %w", err)
		}

		habits = append(habits, newHabit)
		fmt.Printf("✓ New habit '%s' added and marked for today!\n", habitName)
	}

	// Save updated habits
	if err := store.Save(habits); err != nil {
		return fmt.Errorf("failed to save habits: %w", err)
	}

	return nil
}
