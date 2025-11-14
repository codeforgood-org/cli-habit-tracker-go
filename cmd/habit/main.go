// Habit Tracker - A simple CLI tool for tracking daily habits and building streaks.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codeforgood-org/cli-habit-tracker-go/internal/config"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/commands"
	"github.com/codeforgood-org/cli-habit-tracker-go/pkg/storage"
)

const version = "2.0.0"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Parse arguments
	args := os.Args
	if len(args) < 2 {
		printUsage()
		return nil
	}

	// Load configuration
	cfg := config.FromEnv()

	// Initialize storage
	store := storage.NewJSONStorage(cfg.DataFilePath)

	// Parse command
	command := args[1]

	// Route to appropriate handler
	switch command {
	case "list", "ls":
		return commands.List(store)

	case "mark", "done":
		if len(args) < 3 {
			return fmt.Errorf("please provide a habit name")
		}
		habitName := strings.Join(args[2:], " ")
		return commands.Mark(store, habitName)

	case "delete", "del", "rm":
		if len(args) < 3 {
			return fmt.Errorf("please provide a habit name")
		}
		habitName := strings.Join(args[2:], " ")
		return commands.Delete(store, habitName)

	case "reset":
		if len(args) < 3 {
			return fmt.Errorf("please provide a habit name")
		}
		habitName := strings.Join(args[2:], " ")
		return commands.Reset(store, habitName)

	case "stats", "statistics":
		return commands.Stats(store)

	case "export":
		if len(args) < 4 {
			return fmt.Errorf("usage: habit export <format> <output-file>\n  formats: csv, json")
		}
		format := args[2]
		outputPath := args[3]
		return commands.Export(store, format, outputPath)

	case "import":
		if len(args) < 4 {
			return fmt.Errorf("usage: habit import <format> <input-file> [--merge]\n  formats: csv, json")
		}
		format := args[2]
		inputPath := args[3]
		merge := len(args) > 4 && (args[4] == "--merge" || args[4] == "-m")
		return commands.Import(store, format, inputPath, merge)

	case "search", "find":
		if len(args) < 3 {
			return fmt.Errorf("please provide a search query")
		}
		query := strings.Join(args[2:], " ")
		return commands.Search(store, query)

	case "edit", "rename":
		if len(args) < 4 {
			return fmt.Errorf("usage: habit edit <current-name> <new-name>")
		}
		oldName := args[2]
		newName := strings.Join(args[3:], " ")
		return commands.Edit(store, oldName, newName)

	case "backup":
		backupPath := ""
		if len(args) > 2 {
			backupPath = args[2]
		}
		return commands.Backup(store, backupPath)

	case "restore":
		if len(args) < 3 {
			return fmt.Errorf("usage: habit restore <backup-file>")
		}
		backupPath := args[2]
		return commands.Restore(store, backupPath)

	case "version", "-v", "--version":
		fmt.Printf("habit-tracker v%s\n", version)
		return nil

	case "help", "-h", "--help":
		printHelp()
		return nil

	default:
		printUsage()
		return fmt.Errorf("unknown command: %s", command)
	}
}

func printUsage() {
	fmt.Println("Usage: habit <command> [arguments]")
	fmt.Println()
	fmt.Println("Core Commands:")
	fmt.Println("  list              List all habits with their streaks")
	fmt.Println("  mark <name>       Mark a habit as done for today")
	fmt.Println("  delete <name>     Delete a habit")
	fmt.Println("  reset <name>      Reset a habit's streak")
	fmt.Println("  stats             Show habit statistics")
	fmt.Println()
	fmt.Println("Advanced Commands:")
	fmt.Println("  search <query>    Search for habits by name")
	fmt.Println("  edit <old> <new>  Rename a habit")
	fmt.Println("  export <fmt> <f>  Export habits (csv, json)")
	fmt.Println("  import <fmt> <f>  Import habits (csv, json)")
	fmt.Println("  backup [file]     Backup habits data")
	fmt.Println("  restore <file>    Restore from backup")
	fmt.Println()
	fmt.Println("Other:")
	fmt.Println("  version           Show version information")
	fmt.Println("  help              Show detailed help")
	fmt.Println()
	fmt.Println("For more information, run: habit help")
}

func printHelp() {
	fmt.Println("Habit Tracker - Build and maintain daily habits")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  habit <command> [arguments]")
	fmt.Println()
	fmt.Println("CORE COMMANDS:")
	fmt.Println("  list, ls")
	fmt.Println("      List all tracked habits with their current streaks and last completion dates.")
	fmt.Println()
	fmt.Println("  mark <habit-name>, done <habit-name>")
	fmt.Println("      Mark a habit as completed for today. If the habit is new, it will be created.")
	fmt.Println("      Streaks increment when you complete a habit on consecutive days.")
	fmt.Println()
	fmt.Println("  delete <habit-name>, del <habit-name>, rm <habit-name>")
	fmt.Println("      Permanently delete a habit from tracking.")
	fmt.Println()
	fmt.Println("  reset <habit-name>")
	fmt.Println("      Reset a habit's streak to zero and clear its completion date.")
	fmt.Println()
	fmt.Println("  stats, statistics")
	fmt.Println("      Display statistics about all your habits (total, streaks, completion rate).")
	fmt.Println()
	fmt.Println("ADVANCED COMMANDS:")
	fmt.Println("  search <query>, find <query>")
	fmt.Println("      Search for habits by name (case-insensitive substring match).")
	fmt.Println()
	fmt.Println("  edit <current-name> <new-name>, rename <current-name> <new-name>")
	fmt.Println("      Rename an existing habit.")
	fmt.Println()
	fmt.Println("  export <format> <output-file>")
	fmt.Println("      Export habits to a file. Supported formats: csv, json")
	fmt.Println()
	fmt.Println("  import <format> <input-file> [--merge]")
	fmt.Println("      Import habits from a file. Use --merge to merge with existing habits.")
	fmt.Println("      Without --merge, existing habits will be replaced. Supported formats: csv, json")
	fmt.Println()
	fmt.Println("  backup [output-file]")
	fmt.Println("      Create a backup of your habits data. If no file specified, uses timestamp.")
	fmt.Println()
	fmt.Println("  restore <backup-file>")
	fmt.Println("      Restore habits from a backup file. Current data is auto-backed up first.")
	fmt.Println()
	fmt.Println("OTHER COMMANDS:")
	fmt.Println("  version, -v, --version")
	fmt.Println("      Display the version number.")
	fmt.Println()
	fmt.Println("  help, -h, --help")
	fmt.Println("      Display this help message.")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  # Basic usage")
	fmt.Println("  habit mark \"Morning Exercise\"")
	fmt.Println("  habit list")
	fmt.Println("  habit stats")
	fmt.Println("  habit delete \"Old Habit\"")
	fmt.Println()
	fmt.Println("  # Advanced usage")
	fmt.Println("  habit search exercise")
	fmt.Println("  habit edit \"Excercise\" \"Exercise\"")
	fmt.Println("  habit export csv habits.csv")
	fmt.Println("  habit import json habits-backup.json --merge")
	fmt.Println("  habit backup")
	fmt.Println("  habit restore habits-backup-20250113.json")
	fmt.Println()
	fmt.Println("CONFIGURATION:")
	fmt.Println("  Data file location can be customized using the HABIT_DATA_FILE environment variable.")
	fmt.Println("  Default: ~/.habit-tracker/habits.json")
	fmt.Println()
	fmt.Println("  Example:")
	fmt.Println("    export HABIT_DATA_FILE=~/my-habits.json")
	fmt.Println()
}
