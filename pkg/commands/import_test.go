package commands

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

func TestImport_JSON(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	// Create import file
	importPath := filepath.Join(tmpDir, "import.json")
	habits := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
		{Name: "Reading", LastDone: "2025-01-14", Streak: 3},
	}
	data, _ := json.Marshal(habits)
	if err := os.WriteFile(importPath, data, 0644); err != nil {
		t.Fatalf("Failed to create import file: %v", err)
	}

	// Import
	err := Import(store, "json", importPath, false)
	if err != nil {
		t.Fatalf("Import failed: %v", err)
	}

	// Verify
	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Failed to load habits: %v", err)
	}

	if len(loaded) != 2 {
		t.Errorf("Expected 2 habits, got %d", len(loaded))
	}
}

func TestImport_CSV(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	// Create CSV import file
	importPath := filepath.Join(tmpDir, "import.csv")
	file, _ := os.Create(importPath)
	writer := csv.NewWriter(file)
	writer.Write([]string{"Name", "Last Done", "Streak"})
	writer.Write([]string{"Exercise", "2025-01-15", "5"})
	writer.Write([]string{"Reading", "2025-01-14", "3"})
	writer.Flush()
	file.Close()

	// Import
	err := Import(store, "csv", importPath, false)
	if err != nil {
		t.Fatalf("Import failed: %v", err)
	}

	// Verify
	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Failed to load habits: %v", err)
	}

	if len(loaded) != 2 {
		t.Errorf("Expected 2 habits, got %d", len(loaded))
	}
}

func TestImport_MergeMode(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	// Create existing habits
	existing := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-10", Streak: 2},
		{Name: "Meditation", LastDone: "2025-01-15", Streak: 7},
	}
	store.Save(existing)

	// Create import file with overlapping habit
	importPath := filepath.Join(tmpDir, "import.json")
	imported := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
		{Name: "Reading", LastDone: "2025-01-14", Streak: 3},
	}
	data, _ := json.Marshal(imported)
	os.WriteFile(importPath, data, 0644)

	// Import with merge
	err := Import(store, "json", importPath, true)
	if err != nil {
		t.Fatalf("Import failed: %v", err)
	}

	// Verify
	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("Failed to load habits: %v", err)
	}

	// Should have 3 habits: Exercise (updated), Meditation (kept), Reading (added)
	if len(loaded) != 3 {
		t.Errorf("Expected 3 habits after merge, got %d", len(loaded))
	}

	// Verify Exercise was updated
	exercise, _ := loaded.Find("Exercise")
	if exercise == nil {
		t.Fatal("Exercise habit not found")
	}
	if exercise.Streak != 5 {
		t.Errorf("Exercise streak should be 5 (updated), got %d", exercise.Streak)
	}
}

func TestImport_FileNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	err := Import(store, "json", filepath.Join(tmpDir, "nonexistent.json"), false)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestImport_InvalidFormat(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	err := Import(store, "xml", filepath.Join(tmpDir, "import.xml"), false)
	if err == nil {
		t.Error("Expected error for invalid format, got nil")
	}
}
