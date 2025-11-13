# Habit Tracker CLI

A simple, elegant command-line tool for tracking daily habits and building streaks. Built with Go.

## Features

- Track multiple habits with daily completion
- Automatic streak calculation for consecutive days
- Persistent storage in JSON format
- Simple and intuitive CLI interface
- Configurable data file location
- Comprehensive statistics
- Cross-platform support

## Installation

### From Source

```bash
git clone https://github.com/codeforgood-org/cli-habit-tracker-go.git
cd cli-habit-tracker-go
make install
```

Or manually:

```bash
go install github.com/codeforgood-org/cli-habit-tracker-go/cmd/habit@latest
```

### Build from Source

```bash
git clone https://github.com/codeforgood-org/cli-habit-tracker-go.git
cd cli-habit-tracker-go
go build -o habit ./cmd/habit
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

# Delete a habit
habit delete "Old Habit"

# Reset a habit's streak
habit reset "Exercise"

# Get help
habit help
```

### Commands

#### `list` (or `ls`)
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

#### `mark <habit-name>` (or `done`)
Mark a habit as completed for today. Creates the habit if it doesn't exist.

```bash
habit mark "Morning Exercise"
habit mark Reading
habit done "Drink Water"
```

Streak behavior:
- If completed on consecutive days: streak increments
- If there's a gap: streak resets to 1
- If already marked today: shows a message and doesn't change the streak

#### `delete <habit-name>` (or `del`, `rm`)
Permanently remove a habit from tracking.

```bash
habit delete "Old Habit"
habit rm Exercise
```

#### `reset <habit-name>`
Reset a habit's streak to zero and clear its completion date.

```bash
habit reset "Morning Exercise"
```

#### `stats` (or `statistics`)
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

#### `version`
Display the version number.

```bash
habit version
```

#### `help`
Display detailed help information.

```bash
habit help
```

## Configuration

### Data File Location

By default, habits are stored in `~/.habit-tracker/habits.json`. You can customize this location using the `HABIT_DATA_FILE` environment variable:

```bash
export HABIT_DATA_FILE=~/my-habits.json
habit list
```

Or set it for a single command:

```bash
HABIT_DATA_FILE=/tmp/habits.json habit list
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
├── cmd/
│   └── habit/           # Main application entry point
│       └── main.go
├── pkg/
│   ├── models/          # Data models
│   │   ├── habit.go
│   │   └── habit_test.go
│   ├── storage/         # Data persistence
│   │   ├── json.go
│   │   └── json_test.go
│   └── commands/        # CLI command handlers
│       ├── list.go
│       ├── mark.go
│       ├── delete.go
│       ├── reset.go
│       └── stats.go
├── internal/
│   └── config/          # Configuration management
│       └── config.go
├── .github/
│   └── workflows/       # CI/CD workflows
│       └── ci.yml
├── docs/                # Additional documentation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for [Code for Good](https://github.com/codeforgood-org)
- Inspired by the need for simple, effective habit tracking

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/codeforgood-org/cli-habit-tracker-go/issues) page
2. Create a new issue if your problem isn't already reported
3. Include as much detail as possible (OS, Go version, error messages)

## Roadmap

Future enhancements we're considering:

- [ ] Export habits to CSV/PDF
- [ ] Habit categories and tags
- [ ] Reminder notifications
- [ ] Cloud sync support
- [ ] Weekly/monthly reports
- [ ] Habit templates
- [ ] Goal setting (e.g., "Complete 30 days in a row")
- [ ] Charts and visualizations

## Examples

### Daily Routine

```bash
# Morning routine
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

# View all habits
habit list
```

### Manage Habits

```bash
# Remove a habit you no longer want to track
habit delete "Old Habit"

# Start fresh with a habit
habit reset "Exercise"

# Check overall statistics
habit stats
```
