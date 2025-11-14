package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
)

func TestJSONStorage_SaveAndLoad(t *testing.T) {
	// Create temporary directory for test
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_habits.json")

	store := NewJSONStorage(testFile)

	// Create test data
	habits := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
		{Name: "Reading", LastDone: "2025-01-14", Streak: 3},
	}

	// Save habits
	err := store.Save(habits)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Verify file exists
	if !store.Exists() {
		t.Error("Expected file to exist after Save()")
	}

	// Load habits
	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Verify data
	if len(loaded) != len(habits) {
		t.Errorf("Load() length = %v, want %v", len(loaded), len(habits))
	}

	for i, h := range loaded {
		if h.Name != habits[i].Name {
			t.Errorf("Load()[%d].Name = %v, want %v", i, h.Name, habits[i].Name)
		}
		if h.Streak != habits[i].Streak {
			t.Errorf("Load()[%d].Streak = %v, want %v", i, h.Streak, habits[i].Streak)
		}
		if h.LastDone != habits[i].LastDone {
			t.Errorf("Load()[%d].LastDone = %v, want %v", i, h.LastDone, habits[i].LastDone)
		}
	}
}

func TestJSONStorage_LoadNonExistentFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "nonexistent.json")

	store := NewJSONStorage(testFile)

	// Load from non-existent file should return empty list
	habits, err := store.Load()
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if len(habits) != 0 {
		t.Errorf("Load() length = %v, want 0", len(habits))
	}
}

func TestJSONStorage_LoadEmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "empty.json")

	// Create empty file
	err := os.WriteFile(testFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}

	store := NewJSONStorage(testFile)

	// Load from empty file should return empty list
	habits, err := store.Load()
	if err != nil {
		t.Fatalf("Load() error = %v, want nil", err)
	}

	if len(habits) != 0 {
		t.Errorf("Load() length = %v, want 0", len(habits))
	}
}

func TestJSONStorage_Delete(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "delete_test.json")

	store := NewJSONStorage(testFile)

	// Create file with data
	habits := models.HabitList{
		{Name: "Exercise", Streak: 5},
	}
	err := store.Save(habits)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Verify file exists
	if !store.Exists() {
		t.Fatal("Expected file to exist before Delete()")
	}

	// Delete file
	err = store.Delete()
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	// Verify file doesn't exist
	if store.Exists() {
		t.Error("Expected file not to exist after Delete()")
	}
}

func TestJSONStorage_GetPath(t *testing.T) {
	path := "/tmp/test.json"
	store := NewJSONStorage(path)

	if store.GetPath() != path {
		t.Errorf("GetPath() = %v, want %v", store.GetPath(), path)
	}
}

func TestJSONStorage_SaveCreatesDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "subdir", "habits.json")

	store := NewJSONStorage(testFile)

	// Save should create the directory
	habits := models.HabitList{
		{Name: "Exercise", Streak: 5},
	}
	err := store.Save(habits)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Verify file exists
	if !store.Exists() {
		t.Error("Expected file to exist after Save()")
	}
}
