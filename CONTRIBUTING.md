# Contributing to Habit Tracker CLI

Thank you for your interest in contributing to Habit Tracker CLI! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs actual behavior
- **Environment details** (OS, Go version, etc.)
- **Error messages or logs** if applicable

Example:
```
Title: Streak not incrementing on consecutive days

Description:
When marking a habit on consecutive days, the streak remains at 1 instead of incrementing.

Steps to reproduce:
1. Mark habit "Exercise" on 2025-01-14
2. Mark same habit on 2025-01-15
3. Run `habit list`

Expected: Streak should be 2
Actual: Streak remains at 1

Environment:
- OS: Ubuntu 22.04
- Go version: 1.22
- Habit tracker version: 1.0.0
```

### Suggesting Enhancements

Enhancement suggestions are welcome! Please provide:

- **Clear use case** - What problem does it solve?
- **Proposed solution** - How should it work?
- **Alternatives considered** - What other approaches did you think about?
- **Additional context** - Screenshots, mockups, or examples

### Pull Requests

1. **Fork the repository** and create your branch from `main`:
   ```bash
   git checkout -b feature/amazing-feature
   ```

2. **Make your changes**:
   - Write clear, readable code
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed

3. **Test your changes**:
   ```bash
   make test
   make lint
   ```

4. **Commit your changes**:
   ```bash
   git commit -m "Add amazing feature"
   ```

   Use clear, descriptive commit messages following this format:
   - `feat: Add new feature`
   - `fix: Fix bug in streak calculation`
   - `docs: Update README`
   - `test: Add tests for storage`
   - `refactor: Improve code structure`

5. **Push to your fork**:
   ```bash
   git push origin feature/amazing-feature
   ```

6. **Open a Pull Request** with:
   - Clear title and description
   - Reference any related issues
   - Screenshots or examples if applicable

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional but recommended)

### Getting Started

1. **Clone your fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/cli-habit-tracker-go.git
   cd cli-habit-tracker-go
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the project**:
   ```bash
   make build
   ```

4. **Run tests**:
   ```bash
   make test
   ```

## Coding Guidelines

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting (run `make fmt`)
- Run `go vet` before committing (run `make vet`)
- Use meaningful variable and function names
- Add comments for exported functions and types

### Code Organization

```
cli-habit-tracker-go/
├── cmd/habit/          # Application entry point
├── pkg/                # Public packages
│   ├── models/         # Data structures
│   ├── storage/        # Data persistence
│   └── commands/       # Command handlers
└── internal/           # Private packages
    └── config/         # Configuration
```

### Testing

- Write unit tests for all new functionality
- Aim for high code coverage (>80%)
- Use table-driven tests where appropriate
- Test edge cases and error conditions

Example test:
```go
func TestHabit_UpdateStreak(t *testing.T) {
    tests := []struct {
        name    string
        habit   Habit
        want    int
        wantErr bool
    }{
        {
            name:  "consecutive day increments",
            habit: Habit{LastDone: "2025-01-14", Streak: 5},
            want:  6,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Documentation

- Update README.md for user-facing changes
- Add godoc comments for exported functions
- Update CONTRIBUTING.md if changing development process
- Add examples for new features

## Project Structure

### Packages

- **cmd/habit**: Main application entry point
- **pkg/models**: Core data structures (Habit, HabitList)
- **pkg/storage**: Data persistence (JSON storage)
- **pkg/commands**: CLI command implementations
- **internal/config**: Configuration management

### Adding New Commands

1. Create handler in `pkg/commands/`:
   ```go
   // pkg/commands/mycommand.go
   package commands

   func MyCommand(store storage.Storage, args ...string) error {
       // Implementation
   }
   ```

2. Add command to main.go:
   ```go
   case "mycommand":
       return commands.MyCommand(store, args[2:]...)
   ```

3. Add tests in `pkg/commands/mycommand_test.go`

4. Update help text and README

## Release Process

Maintainers will handle releases following semantic versioning:

- **Major** (1.0.0): Breaking changes
- **Minor** (0.1.0): New features (backward compatible)
- **Patch** (0.0.1): Bug fixes

## Questions?

- Open an issue for questions
- Check existing issues and pull requests
- Read the README.md and documentation

## Recognition

Contributors will be recognized in:
- GitHub contributors list
- Release notes
- Project documentation

Thank you for contributing to Habit Tracker CLI!
