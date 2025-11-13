// Package commands implements CLI command handlers.
package commands

import (
	"fmt"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// List displays all tracked habits with their streaks.
func List(store storage.Storage) error {
	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	if len(habits) == 0 {
		fmt.Println("No habits tracked.")
		fmt.Println("\nTo start tracking a habit, use:")
		fmt.Println("  habit mark <habit-name>")
		return nil
	}

	fmt.Printf("📋 Tracking %d habit(s):\n\n", len(habits))
	for _, h := range habits {
		fmt.Println(h.String())
	}

	return nil
}
