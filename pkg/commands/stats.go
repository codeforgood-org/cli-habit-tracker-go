// Package commands implements CLI command handlers.
package commands

import (
	"fmt"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Stats displays statistics about all habits.
func Stats(store storage.Storage) error {
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	if len(habits) == 0 {
		fmt.Println("No habits tracked yet.")
		return nil
	}

	stats := habits.Stats()

	fmt.Println("📊 Habit Statistics:")
	fmt.Println()
	fmt.Printf("  Total habits:       %d\n", stats["total"])
	fmt.Printf("  Marked today:       %d\n", stats["marked_today"])
	fmt.Printf("  Longest streak:     %d day(s)\n", stats["max_streak"])
	fmt.Printf("  Total streak days:  %d\n", stats["total_streak"])
	fmt.Printf("  Average streak:     %.1f day(s)\n", stats["avg_streak"])

	return nil
}
