package models

import (
	"testing"
	"time"
)

func TestHabit_Validate(t *testing.T) {
	tests := []struct {
		name    string
		habit   Habit
		wantErr bool
	}{
		{
			name:    "valid habit",
			habit:   Habit{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
			wantErr: false,
		},
		{
			name:    "empty name",
			habit:   Habit{Name: "", LastDone: "2025-01-15", Streak: 5},
			wantErr: true,
		},
		{
			name:    "whitespace name",
			habit:   Habit{Name: "   ", LastDone: "2025-01-15", Streak: 5},
			wantErr: true,
		},
		{
			name:    "invalid date format",
			habit:   Habit{Name: "Exercise", LastDone: "15-01-2025", Streak: 5},
			wantErr: true,
		},
		{
			name:    "negative streak",
			habit:   Habit{Name: "Exercise", LastDone: "2025-01-15", Streak: -1},
			wantErr: true,
		},
		{
			name:    "zero streak is valid",
			habit:   Habit{Name: "Exercise", LastDone: "", Streak: 0},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.habit.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHabit_UpdateStreak(t *testing.T) {
	tests := []struct {
		name         string
		habit        Habit
		today        time.Time
		wantStreak   int
		wantLastDone string
		wantErr      bool
	}{
		{
			name:         "first time marking",
			habit:        Habit{Name: "Exercise", LastDone: "", Streak: 0},
			today:        time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			wantStreak:   1,
			wantLastDone: "2025-01-15",
			wantErr:      false,
		},
		{
			name:         "consecutive day",
			habit:        Habit{Name: "Exercise", LastDone: "2025-01-14", Streak: 5},
			today:        time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			wantStreak:   6,
			wantLastDone: "2025-01-15",
			wantErr:      false,
		},
		{
			name:         "gap in days",
			habit:        Habit{Name: "Exercise", LastDone: "2025-01-10", Streak: 5},
			today:        time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			wantStreak:   1,
			wantLastDone: "2025-01-15",
			wantErr:      false,
		},
		{
			name:         "already marked today",
			habit:        Habit{Name: "Exercise", LastDone: "2025-01-15", Streak: 5},
			today:        time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			wantStreak:   5,
			wantLastDone: "2025-01-15",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.habit.UpdateStreak(tt.today)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStreak() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.habit.Streak != tt.wantStreak {
					t.Errorf("UpdateStreak() streak = %v, want %v", tt.habit.Streak, tt.wantStreak)
				}
				if tt.habit.LastDone != tt.wantLastDone {
					t.Errorf("UpdateStreak() lastDone = %v, want %v", tt.habit.LastDone, tt.wantLastDone)
				}
			}
		})
	}
}

func TestHabit_IsMarkedToday(t *testing.T) {
	today := time.Date(2025, 1, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name   string
		habit  Habit
		date   time.Time
		want   bool
	}{
		{
			name:   "marked today",
			habit:  Habit{Name: "Exercise", LastDone: "2025-01-15"},
			date:   today,
			want:   true,
		},
		{
			name:   "not marked today",
			habit:  Habit{Name: "Exercise", LastDone: "2025-01-14"},
			date:   today,
			want:   false,
		},
		{
			name:   "never marked",
			habit:  Habit{Name: "Exercise", LastDone: ""},
			date:   today,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.habit.IsMarkedToday(tt.date); got != tt.want {
				t.Errorf("IsMarkedToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHabitList_Find(t *testing.T) {
	habits := HabitList{
		{Name: "Exercise", Streak: 5},
		{Name: "Reading", Streak: 3},
		{Name: "Meditation", Streak: 10},
	}

	tests := []struct {
		name      string
		searchFor string
		wantFound bool
		wantIndex int
	}{
		{
			name:      "exact match",
			searchFor: "Exercise",
			wantFound: true,
			wantIndex: 0,
		},
		{
			name:      "case insensitive",
			searchFor: "READING",
			wantFound: true,
			wantIndex: 1,
		},
		{
			name:      "not found",
			searchFor: "Swimming",
			wantFound: false,
			wantIndex: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			habit, index := habits.Find(tt.searchFor)
			if tt.wantFound && habit == nil {
				t.Error("Find() expected to find habit but got nil")
			}
			if !tt.wantFound && habit != nil {
				t.Error("Find() expected not to find habit but got one")
			}
			if index != tt.wantIndex {
				t.Errorf("Find() index = %v, want %v", index, tt.wantIndex)
			}
		})
	}
}

func TestHabitList_Remove(t *testing.T) {
	tests := []struct {
		name      string
		habits    HabitList
		index     int
		wantErr   bool
		wantLen   int
	}{
		{
			name: "remove first",
			habits: HabitList{
				{Name: "Exercise"},
				{Name: "Reading"},
			},
			index:   0,
			wantErr: false,
			wantLen: 1,
		},
		{
			name: "remove last",
			habits: HabitList{
				{Name: "Exercise"},
				{Name: "Reading"},
			},
			index:   1,
			wantErr: false,
			wantLen: 1,
		},
		{
			name: "index out of range",
			habits: HabitList{
				{Name: "Exercise"},
			},
			index:   5,
			wantErr: true,
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.habits.Remove(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(tt.habits) != tt.wantLen {
				t.Errorf("Remove() length = %v, want %v", len(tt.habits), tt.wantLen)
			}
		})
	}
}

func TestHabitList_Stats(t *testing.T) {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	tests := []struct {
		name   string
		habits HabitList
		want   map[string]interface{}
	}{
		{
			name:   "empty list",
			habits: HabitList{},
			want: map[string]interface{}{
				"total":        0,
				"max_streak":   0,
				"total_streak": 0,
				"avg_streak":   0.0,
			},
		},
		{
			name: "multiple habits",
			habits: HabitList{
				{Name: "Exercise", Streak: 5, LastDone: today.Format("2006-01-02")},
				{Name: "Reading", Streak: 10, LastDone: yesterday.Format("2006-01-02")},
				{Name: "Meditation", Streak: 3, LastDone: today.Format("2006-01-02")},
			},
			want: map[string]interface{}{
				"total":        3,
				"max_streak":   10,
				"total_streak": 18,
				"avg_streak":   6.0,
				"marked_today": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.habits.Stats()
			for key, wantVal := range tt.want {
				gotVal, ok := got[key]
				if !ok {
					t.Errorf("Stats() missing key %s", key)
					continue
				}
				if gotVal != wantVal {
					t.Errorf("Stats()[%s] = %v, want %v", key, gotVal, wantVal)
				}
			}
		})
	}
}
