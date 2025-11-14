package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/models"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

func TestExport_CSV(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	// Create test data
	habits := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
		{Name: "Reading", LastDone: "2025-01-14", Streak: 3},
	}
	if err := store.Save(habits); err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}

	// Export to CSV
	outputPath := filepath.Join(tmpDir, "export.csv")
	err := Export(store, "csv", outputPath)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("Export file was not created")
	}

	// Verify content
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read export file: %v", err)
	}

	content := string(data)
	if !contains(content, "Exercise") || !contains(content, "Reading") {
		t.Error("Export file missing expected habit data")
	}
}

func TestExport_JSON(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	habits := models.HabitList{
		{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
	}
	if err := store.Save(habits); err != nil {
		t.Fatalf("Failed to save test data: %v", err)
	}

	outputPath := filepath.Join(tmpDir, "export.json")
	err := Export(store, "json", outputPath)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("Export file was not created")
	}
}

func TestExport_InvalidFormat(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	err := Export(store, "xml", filepath.Join(tmpDir, "export.xml"))
	if err == nil {
		t.Error("Expected error for invalid format, got nil")
	}
}

func TestExport_EmptyHabits(t *testing.T) {
	tmpDir := t.TempDir()
	storePath := filepath.Join(tmpDir, "habits.json")
	store := storage.NewJSONStorage(storePath)

	err := Export(store, "csv", filepath.Join(tmpDir, "export.csv"))
	if err == nil {
		t.Error("Expected error for empty habits, got nil")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
