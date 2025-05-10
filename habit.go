package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const dataFile = "habits.json"

type Habit struct {
	Name       string    `json:"name"`
	LastDone   string    `json:"last_done"`
	Streak     int       `json:"streak"`
}

func loadHabits() ([]Habit, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return []Habit{}, nil
	}
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	var habits []Habit
	err = json.Unmarshal(data, &habits)
	return habits, err
}

func saveHabits(habits []Habit) error {
	data, err := json.MarshalIndent(habits, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func listHabits(habits []Habit) {
	if len(habits) == 0 {
		fmt.Println("No habits tracked.")
		return
	}
	for _, h := range habits {
		fmt.Printf("- %s | Streak: %d | Last done: %s\n", h.Name, h.Streak, h.LastDone)
	}
}

func markHabit(habits []Habit, name string) []Habit {
	now := time.Now().Format("2006-01-02")
	for i, h := range habits {
		if strings.EqualFold(h.Name, name) {
			if h.LastDone == now {
				fmt.Println("Already marked today.")
				return habits
			}
			last, _ := time.Parse("2006-01-02", h.LastDone)
			today := time.Now()
			diff := int(today.Sub(last).Hours() / 24)
			if diff == 1 {
				h.Streak++
			} else {
				h.Streak = 1
			}
			h.LastDone = now
			habits[i] = h
			fmt.Println("Marked habit as done today.")
			return habits
		}
	}
	// new habit
	habits = append(habits, Habit{Name: name, LastDone: now, Streak: 1})
	fmt.Println("New habit added and marked for today.")
	return habits
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  habit list")
	fmt.Println("  habit mark <habit-name>")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
		return
	}
	command := args[1]
	habits, err := loadHabits()
	if err != nil {
		fmt.Println("Error loading habits:", err)
		return
	}

	switch command {
	case "list":
		listHabits(habits)
	case "mark":
		if len(args) < 3 {
			fmt.Println("Please provide a habit name.")
			return
		}
		name := strings.Join(args[2:], " ")
		habits = markHabit(habits, name)
		if err := saveHabits(habits); err != nil {
			fmt.Println("Error saving habits:", err)
		}
	default:
		usage()
	}
}
