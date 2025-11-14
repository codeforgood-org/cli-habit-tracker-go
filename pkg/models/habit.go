// Package models defines the core data structures for the habit tracker.
package models

import (
	"fmt"
	"strings"
	"time"
)

// Habit represents a single habit being tracked with its streak information.
type Habit struct {
	Name     string `json:"name"`     // Name of the habit
	LastDone string `json:"last_done"` // Last completion date in YYYY-MM-DD format
	Streak   int    `json:"streak"`   // Current streak count (consecutive days)
}

// Validate checks if the habit has valid data.
func (h *Habit) Validate() error {
	if strings.TrimSpace(h.Name) == "" {
		return fmt.Errorf("habit name cannot be empty")
	}
	if h.LastDone != "" {
		_, err := time.Parse("2006-01-02", h.LastDone)
		if err != nil {
			return fmt.Errorf("invalid date format for LastDone: %w", err)
		}
	}
	if h.Streak < 0 {
		return fmt.Errorf("streak cannot be negative")
	}
	return nil
}

// UpdateStreak updates the habit's streak based on the last completion date.
// It increments the streak if completed consecutively, otherwise resets to 1.
func (h *Habit) UpdateStreak(today time.Time) error {
	todayStr := today.Format("2006-01-02")

	// Check if already marked today
	if h.LastDone == todayStr {
		return fmt.Errorf("habit already marked for today")
	}

	// If this is a new habit or first time marking
	if h.LastDone == "" {
		h.Streak = 1
		h.LastDone = todayStr
		return nil
	}

	// Parse the last done date
	lastDone, err := time.Parse("2006-01-02", h.LastDone)
	if err != nil {
		return fmt.Errorf("invalid last done date: %w", err)
	}

	// Calculate day difference
	diff := int(today.Sub(lastDone).Hours() / 24)

	// Update streak based on difference
	if diff == 1 {
		// Consecutive day - increment streak
		h.Streak++
	} else {
		// Gap in days - reset streak
		h.Streak = 1
	}

	h.LastDone = todayStr
	return nil
}

// IsMarkedToday checks if the habit was marked on the given date.
func (h *Habit) IsMarkedToday(today time.Time) bool {
	return h.LastDone == today.Format("2006-01-02")
}

// DaysSinceLastDone returns the number of days since the habit was last completed.
func (h *Habit) DaysSinceLastDone() (int, error) {
	if h.LastDone == "" {
		return -1, fmt.Errorf("habit has never been completed")
	}

	lastDone, err := time.Parse("2006-01-02", h.LastDone)
	if err != nil {
		return -1, fmt.Errorf("invalid last done date: %w", err)
	}

	today := time.Now()
	diff := int(today.Sub(lastDone).Hours() / 24)
	return diff, nil
}

// String returns a formatted string representation of the habit.
func (h *Habit) String() string {
	lastDone := h.LastDone
	if lastDone == "" {
		lastDone = "Never"
	}
	return fmt.Sprintf("- %s | Streak: %d | Last done: %s", h.Name, h.Streak, lastDone)
}

// HabitList represents a collection of habits.
type HabitList []Habit

// Find returns a habit by name (case-insensitive).
func (hl HabitList) Find(name string) (*Habit, int) {
	for i, h := range hl {
		if strings.EqualFold(h.Name, name) {
			return &hl[i], i
		}
	}
	return nil, -1
}

// Contains checks if a habit with the given name exists.
func (hl HabitList) Contains(name string) bool {
	habit, _ := hl.Find(name)
	return habit != nil
}

// Remove removes a habit from the list by index.
func (hl *HabitList) Remove(index int) error {
	if index < 0 || index >= len(*hl) {
		return fmt.Errorf("index out of range")
	}
	*hl = append((*hl)[:index], (*hl)[index+1:]...)
	return nil
}

// Add adds a new habit to the list.
func (hl *HabitList) Add(habit Habit) error {
	if err := habit.Validate(); err != nil {
		return err
	}
	*hl = append(*hl, habit)
	return nil
}

// Stats returns statistics about the habit list.
func (hl HabitList) Stats() map[string]interface{} {
	if len(hl) == 0 {
		return map[string]interface{}{
			"total":        0,
			"max_streak":   0,
			"total_streak": 0,
			"avg_streak":   0.0,
		}
	}

	maxStreak := 0
	totalStreak := 0
	markedToday := 0
	today := time.Now()

	for _, h := range hl {
		if h.Streak > maxStreak {
			maxStreak = h.Streak
		}
		totalStreak += h.Streak
		if h.IsMarkedToday(today) {
			markedToday++
		}
	}

	return map[string]interface{}{
		"total":        len(hl),
		"max_streak":   maxStreak,
		"total_streak": totalStreak,
		"avg_streak":   float64(totalStreak) / float64(len(hl)),
		"marked_today": markedToday,
	}
}
