package commands

import (
	"path/filepath"
	"testing"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

func TestSearch(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	habits := models.HabitList{
		{Name: "Morning Exercise", LastDone: "2025-01-15", Streak: 5},
		{Name: "Evening Reading", LastDone: "2025-01-14", Streak: 3},
		{Name: "Meditation", LastDone: "2025-01-15", Streak: 10},
	}
	store.Save(habits)

	// Test case-insensitive search
	err := Search(store, "exercise")
	if err != nil {
		t.Errorf("Search failed: %v", err)
	}

	// Test partial match
	err = Search(store, "read")
	if err != nil {
		t.Errorf("Search failed: %v", err)
	}
}

func TestSearch_EmptyQuery(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	err := Search(store, "")
	if err == nil {
		t.Error("Expected error for empty query, got nil")
	}
}

func TestSearch_NoResults(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	habits := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
	}
	store.Save(habits)

	// Should not error, just show no results
	err := Search(store, "nonexistent")
	if err != nil {
		t.Errorf("Search should not error on no results: %v", err)
	}
}
