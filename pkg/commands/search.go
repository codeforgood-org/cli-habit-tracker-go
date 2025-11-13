// Package commands implements CLI command handlers.
package commands

import (
	"fmt"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

// Search searches for habits matching a query string.
func Search(store storage.Storage, query string) error {
	query = strings.TrimSpace(query)
	if query == "" {
		return fmt.Errorf("search query cannot be empty")
	}

	habits, err := store.Load()
	if err != nil {
		return fmt.Errorf("failed to load habits: %w", err)
	}

	if len(habits) == 0 {
		fmt.Println("No habits tracked.")
		return nil
	}

	// Search for matching habits (case-insensitive substring match)
	queryLower := strings.ToLower(query)
	var matches []string
	matchCount := 0

	for _, habit := range habits {
		if strings.Contains(strings.ToLower(habit.Name), queryLower) {
			matches = append(matches, habit.String())
			matchCount++
		}
	}

	if matchCount == 0 {
		fmt.Printf("No habits found matching '%s'\n", query)
		return nil
	}

	fmt.Printf("Found %d habit(s) matching '%s':\n\n", matchCount, query)
	for _, match := range matches {
		fmt.Println(match)
	}

	return nil
}
