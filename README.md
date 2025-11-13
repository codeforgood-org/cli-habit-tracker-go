# Habit Tracker CLI

[![CI](https://github.com/codeforgood-org/cli-habit-tracker-go/workflows/CI/badge.svg)](https://github.com/codeforgood-org/cli-habit-tracker-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/codeforgood-org/cli-habit-tracker-go)](https://goreportcard.com/report/github.com/codeforgood-org/cli-habit-tracker-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/codeforgood-org/cli-habit-tracker-go)](https://github.com/codeforgood-org/cli-habit-tracker-go/releases/latest)
[![Go Version](https://img.shields.io/github/go-mod/go-version/codeforgood-org/cli-habit-tracker-go)](go.mod)

A simple, elegant command-line tool for tracking daily habits and building streaks. Built with Go.

## Features

### Core Features
- ✅ Track unlimited habits with daily completion
- 📈 Automatic streak calculation for consecutive days
- 💾 Persistent storage in JSON format
- 🎯 Simple and intuitive CLI interface
- ⚙️ Configurable data file location
- 📊 Comprehensive statistics
- 🌍 Cross-platform support (Linux, macOS, Windows)

### Advanced Features
- 🔍 **Search** habits by name
- ✏️ **Rename** habits
- 📤 **Export** to CSV or JSON
- 📥 **Import** from CSV or JSON with merge support
- 💾 **Backup & Restore** functionality
- 🚀 **Shell completions** (Bash, Zsh, Fish)
- 🐳 **Docker support**

## Installation

### Quick Install (Recommended)

```bash
curl -sfL https://raw.githubusercontent.com/codeforgood-org/cli-habit-tracker-go/main/install.sh | sh
```

### Using Go

```bash
go install github.com/codeforgood-org/cli-habit-tracker-go/cmd/habit@latest
```

### Using Homebrew (macOS/Linux)

```bash
brew tap codeforgood-org/tap
brew install habit-tracker
```

### Download Binary

Download the latest binary for your platform from the [releases page](https://github.com/codeforgood-org/cli-habit-tracker-go/releases).

### Build from Source

```bash
git clone https://github.com/codeforgood-org/cli-habit-tracker-go.git
cd cli-habit-tracker-go
make build
sudo make install
```

### Docker

```bash
docker pull ghcr.io/codeforgood-org/habit-tracker:latest
docker run --rm -v ~/.habit-tracker:/root/.habit-tracker ghcr.io/codeforgood-org/habit-tracker list
```

## Usage

### Quick Start

```bash
# Mark a habit as done for today (creates it if new)
habit mark "Morning Exercise"

# List all habits with their streaks
habit list

# View statistics
habit stats

# Search for habits
habit search exercise

# Delete a habit
habit delete "Old Habit"

# Rename a habit
habit edit "Excercise" "Exercise"

# Export habits to CSV
habit export csv habits.csv

# Import habits from JSON
habit import json backup.json --merge

# Create a backup
habit backup

# Restore from backup
habit restore habits-backup.json
```

### Commands

#### Core Commands

##### `list` (or `ls`)
List all tracked habits with their current streaks and last completion dates.

```bash
habit list
```

Output:
```
📋 Tracking 3 habit(s):

- Morning Exercise | Streak: 7 | Last done: 2025-01-15
- Reading | Streak: 3 | Last done: 2025-01-14
- Meditation | Streak: 10 | Last done: 2025-01-15
```

##### `mark <habit-name>` (or `done`)
Mark a habit as completed for today. Creates the habit if it doesn't exist.

```bash
habit mark "Morning Exercise"
habit done Reading
```

Streak behavior:
- ✅ **Consecutive days**: streak increments
- ⏭️ **Gap in days**: streak resets to 1
- ℹ️ **Already marked**: shows a message, doesn't change streak

##### `delete <habit-name>` (or `del`, `rm`)
Permanently remove a habit from tracking.

```bash
habit delete "Old Habit"
habit rm Exercise
```

##### `reset <habit-name>`
Reset a habit's streak to zero and clear its completion date.

```bash
habit reset "Morning Exercise"
```

##### `stats` (or `statistics`)
Display comprehensive statistics about all your habits.

```bash
habit stats
```

Output:
```
📊 Habit Statistics:

  Total habits:       5
  Marked today:       3
  Longest streak:     15 day(s)
  Total streak days:  42
  Average streak:     8.4 day(s)
```

#### Advanced Commands

##### `search <query>` (or `find`)
Search for habits by name (case-insensitive substring match).

```bash
habit search exercise
habit find read
```

##### `edit <current-name> <new-name>` (or `rename`)
Rename an existing habit.

```bash
habit edit "Excercise" "Exercise"
habit rename "Old Name" "New Name"
```

##### `export <format> <output-file>`
Export habits to a file. Supported formats: `csv`, `json`

```bash
habit export csv habits.csv
habit export json habits-backup.json
```

##### `import <format> <input-file> [--merge]`
Import habits from a file. Use `--merge` to merge with existing habits.

```bash
# Replace existing habits
habit import json habits-backup.json

# Merge with existing habits
habit import csv habits.csv --merge
```

##### `backup [output-file]`
Create a backup of your habits data. If no file specified, uses timestamp.

```bash
habit backup                        # Auto-named backup
habit backup my-backup.json         # Custom filename
```

##### `restore <backup-file>`
Restore habits from a backup file. Current data is auto-backed up first.

```bash
habit restore habits-backup-20250113.json
```

#### Other Commands

##### `version`
Display the version number.

```bash
habit version
```

##### `help`
Display detailed help information.

```bash
habit help
```

## Configuration

### Data File Location

By default, habits are stored in `~/.habit-tracker/habits.json`. You can customize this location using the `HABIT_DATA_FILE` environment variable:

```bash
# Temporary override
HABIT_DATA_FILE=~/my-habits.json habit list

# Permanent override (add to ~/.bashrc or ~/.zshrc)
export HABIT_DATA_FILE=~/my-habits.json
```

### Shell Completions

Enable tab completion for your shell:

**Bash:**
```bash
sudo cp completions/habit.bash /etc/bash_completion.d/habit
source /etc/bash_completion.d/habit
```

**Zsh:**
```bash
mkdir -p ~/.zsh/completions
cp completions/habit.zsh ~/.zsh/completions/_habit
# Add to ~/.zshrc:
fpath=(~/.zsh/completions $fpath)
autoload -U compinit && compinit
```

**Fish:**
```bash
cp completions/habit.fish ~/.config/fish/completions/
```

### Data Format

Habits are stored in JSON format:

```json
[
  {
    "name": "Morning Exercise",
    "last_done": "2025-01-15",
    "streak": 7
  },
  {
    "name": "Reading",
    "last_done": "2025-01-14",
    "streak": 3
  }
]
```

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional but recommended)

### Building

```bash
make build
```

### Running Tests

```bash
make test
```

### Test Coverage

```bash
make coverage
```

### Linting

```bash
make lint
```

### Clean Build Artifacts

```bash
make clean
```

## Project Structure

```
cli-habit-tracker-go/
├── cmd/habit/              # Main application entry point
├── pkg/
│   ├── models/            # Data models with business logic
│   ├── storage/           # JSON persistence layer
│   └── commands/          # CLI command handlers
├── internal/config/       # Configuration management
├── docs/                  # Documentation
│   ├── ARCHITECTURE.md    # Architecture documentation
│   └── FAQ.md            # Frequently asked questions
├── completions/           # Shell completion scripts
├── .github/workflows/     # CI/CD pipelines
├── Dockerfile            # Docker configuration
├── .goreleaser.yml       # Release automation
├── install.sh            # Installation script
├── Makefile             # Build automation
└── [documentation files]
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Documentation

- [Architecture Documentation](docs/ARCHITECTURE.md) - Design and architecture details
- [Contributing Guide](CONTRIBUTING.md) - How to contribute
- [FAQ](docs/FAQ.md) - Frequently asked questions
- [Changelog](CHANGELOG.md) - Version history and changes

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for [Code for Good](https://github.com/codeforgood-org)
- Inspired by the need for simple, effective habit tracking
- Thanks to all [contributors](https://github.com/codeforgood-org/cli-habit-tracker-go/graphs/contributors)

## Support

If you encounter any issues or have questions:

1. Check the [FAQ](docs/FAQ.md)
2. Search [existing issues](https://github.com/codeforgood-org/cli-habit-tracker-go/issues)
3. Create a [new issue](https://github.com/codeforgood-org/cli-habit-tracker-go/issues/new)

## Roadmap

Future enhancements we're considering:

- [ ] Cloud sync support (Dropbox, Google Drive, iCloud)
- [ ] Reminder notifications
- [ ] Charts and visualizations (terminal-based)
- [ ] Goal setting (e.g., "Complete 30 days in a row")
- [ ] Habit categories and tags
- [ ] Weekly/monthly reports
- [ ] Habit templates
- [ ] Mobile companion app
- [ ] Web dashboard
- [ ] API for integrations

## Examples

### Daily Routine

```bash
# Morning routine
habit mark "Wake up at 6am"
habit mark "Morning Exercise"
habit mark "Meditation"
habit mark "Healthy Breakfast"

# Check progress
habit list

# Evening review
habit stats
```

### Track Multiple Habits

```bash
# Add various habits
habit mark "Read 30 minutes"
habit mark "Drink 8 glasses of water"
habit mark "Learn Spanish"
habit mark "Code Review"
habit mark "Journal"

# View all habits
habit list

# Search for specific habits
habit search water
```

### Manage and Organize

```bash
# Fix typos
habit edit "Excercise" "Exercise"

# Remove old habits
habit delete "Deprecated Habit"

# Start fresh with a habit
habit reset "Exercise"

# Check overall statistics
habit stats
```

### Backup and Export

```bash
# Create daily backup
habit backup ~/backups/habits-$(date +%Y%m%d).json

# Export to CSV for analysis
habit export csv habits-export.csv

# Import from another device
habit import json habits-from-laptop.json --merge
```

### Advanced Usage

```bash
# Use different data files for work/personal
alias work-habits='HABIT_DATA_FILE=~/work-habits.json habit'
alias personal-habits='HABIT_DATA_FILE=~/personal-habits.json habit'

work-habits mark "Code Review"
personal-habits mark "Exercise"

# Automated daily reminder script
#!/bin/bash
if habit list | grep -q "Last done: $(date +%Y-%m-%d)"; then
    echo "Great job! You've tracked habits today!"
else
    echo "Don't forget to track your habits!"
    habit list
fi
```

## Star History

If you find this project useful, please consider giving it a star ⭐

## Community

- [GitHub Discussions](https://github.com/codeforgood-org/cli-habit-tracker-go/discussions) - Ask questions, share ideas
- [Issue Tracker](https://github.com/codeforgood-org/cli-habit-tracker-go/issues) - Report bugs, request features

---

Made with ❤️ by the Code for Good community
