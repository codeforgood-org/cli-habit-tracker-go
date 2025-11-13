# Architecture Documentation

## Overview

Habit Tracker CLI is a command-line application built with Go that follows clean architecture principles with clear separation of concerns.

## Design Principles

1. **Separation of Concerns**: Each package has a single, well-defined responsibility
2. **Testability**: All business logic is unit testable
3. **Dependency Injection**: Dependencies are injected rather than hardcoded
4. **Interface-based Design**: Using interfaces for flexibility and testing
5. **Simplicity**: Keep it simple and maintainable

## Architecture Layers

```
┌─────────────────────────────────────┐
│         CLI Interface               │
│      (cmd/habit/main.go)            │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│       Command Handlers              │
│      (pkg/commands/*)               │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│       Business Logic                │
│      (pkg/models/*)                 │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│       Data Storage                  │
│      (pkg/storage/*)                │
└─────────────────────────────────────┘
```

## Package Structure

### cmd/habit

**Purpose**: Application entry point and CLI argument parsing

**Responsibilities**:
- Parse command-line arguments
- Initialize configuration
- Route commands to appropriate handlers
- Handle top-level errors and exit codes

**Key Components**:
- `main()`: Entry point
- `run()`: Main application logic
- Command routing switch statement
- Help and usage text

### pkg/models

**Purpose**: Core domain models and business logic

**Responsibilities**:
- Define data structures (`Habit`, `HabitList`)
- Implement business rules (streak calculation, validation)
- Provide utility methods for habit operations

**Key Components**:
- `Habit`: Individual habit with streak tracking
  - `UpdateStreak()`: Calculates and updates streak
  - `Validate()`: Ensures data integrity
  - `IsMarkedToday()`: Checks completion status
- `HabitList`: Collection of habits
  - `Find()`: Locate habits by name
  - `Stats()`: Calculate statistics
  - `Add()`, `Remove()`: Manage habits

**Business Rules**:
1. Streaks increment only on consecutive days
2. Missing a day resets streak to 1
3. Habit names are case-insensitive
4. Dates are stored in YYYY-MM-DD format

### pkg/storage

**Purpose**: Data persistence layer

**Responsibilities**:
- Save and load habits from persistent storage
- Abstract storage implementation details
- Handle file I/O errors gracefully

**Key Components**:
- `Storage`: Interface defining persistence contract
- `JSONStorage`: JSON file-based implementation
  - `Load()`: Read habits from file
  - `Save()`: Write habits to file
  - `Delete()`: Remove storage file
  - `Exists()`: Check if file exists

**Design Decisions**:
- Uses interface to allow future storage backends (SQLite, PostgreSQL, etc.)
- JSON format for human-readable data
- Automatic directory creation
- Graceful handling of missing files

### pkg/commands

**Purpose**: CLI command implementations

**Responsibilities**:
- Implement specific command logic
- Coordinate between models and storage
- Format output for users
- Handle command-specific errors

**Key Components**:
- `List()`: Display all habits
- `Mark()`: Mark habit as complete
- `Delete()`: Remove a habit
- `Reset()`: Reset habit streak
- `Stats()`: Show statistics

**Command Flow**:
1. Load habits from storage
2. Perform operation on habits
3. Save updated habits (if modified)
4. Display result to user

### internal/config

**Purpose**: Configuration management

**Responsibilities**:
- Manage application settings
- Handle environment variables
- Provide default configurations

**Key Components**:
- `Config`: Configuration structure
- `Default()`: Default configuration
- `FromEnv()`: Environment-based configuration

**Configuration Sources** (in order of precedence):
1. Environment variables
2. Default values

## Data Flow

### Example: Marking a Habit

```
User Input: habit mark "Exercise"
       │
       ▼
cmd/habit/main.go
       │ Parse args
       │ Initialize storage
       ▼
pkg/commands/mark.go
       │ Load habits
       │ Find or create habit
       │ Update streak
       │ Save habits
       ▼
pkg/models/habit.go
       │ UpdateStreak()
       │ Validate()
       ▼
pkg/storage/json.go
       │ Save()
       │ Write to file
       ▼
~/.habit-tracker/habits.json
```

## Error Handling

### Strategy

1. **Bubble up errors**: Errors are returned up the call stack
2. **Context-aware messages**: Add context at each layer
3. **User-friendly output**: Format errors for CLI users
4. **No panics**: Use error returns instead of panics

### Error Flow

```go
// Low-level error
err := os.ReadFile(path)
// -> wrapped with context
return fmt.Errorf("failed to read file: %w", err)
// -> handled in command
return fmt.Errorf("failed to load habits: %w", err)
// -> displayed to user
fmt.Fprintf(os.Stderr, "Error: %v\n", err)
```

## Testing Strategy

### Unit Tests

- **Models**: Test business logic in isolation
- **Storage**: Test with temporary files
- **Commands**: Mock storage interface

### Test Organization

```
pkg/models/
  ├── habit.go
  └── habit_test.go       # Unit tests

pkg/storage/
  ├── json.go
  └── json_test.go        # Integration tests with temp files
```

### Test Patterns

1. **Table-driven tests**: Multiple test cases in one test function
2. **Temporary directories**: Use `t.TempDir()` for file tests
3. **Clear test names**: Descriptive test case names

## Future Enhancements

### Planned Improvements

1. **Additional Storage Backends**
   - SQLite for better querying
   - Cloud storage for sync

2. **Enhanced Features**
   - Habit categories/tags
   - Reminders and notifications
   - Export to CSV/PDF
   - Charts and visualizations

3. **Architecture Changes**
   - Plugin system for storage backends
   - Event-driven architecture for notifications
   - Repository pattern for data access

### Extensibility Points

1. **Storage Interface**: Easy to add new backends
2. **Command Pattern**: Simple to add new commands
3. **Model Methods**: Extend habit functionality

## Performance Considerations

### Current Implementation

- Small data files (typically < 1KB)
- Fast JSON parsing
- In-memory operations
- File I/O only on load/save

### Scaling Considerations

For large numbers of habits (1000+):
- Consider indexing for fast lookups
- Batch operations for efficiency
- Lazy loading of data
- Database backend for querying

## Security

### Current Measures

1. **File Permissions**: 0644 for data files
2. **Input Validation**: Validate all user input
3. **No Shell Execution**: Pure Go implementation
4. **Safe File Paths**: Use `filepath` package

### Future Enhancements

1. **Data Encryption**: Encrypt habits.json
2. **Authentication**: For cloud sync
3. **Input Sanitization**: Additional validation

## Dependencies

### Standard Library Only

The project uses only Go standard library:
- `encoding/json`: Data serialization
- `os`: File operations
- `time`: Date handling
- `fmt`: Formatting
- `strings`: String manipulation

**Benefits**:
- No external dependencies
- Fast builds
- Easy maintenance
- Security (fewer supply chain risks)

## Build and Deployment

### Build Process

```bash
go build -ldflags="-s -w" -o habit ./cmd/habit
```

Flags:
- `-s`: Strip symbol table
- `-w`: Strip debug info
- Reduces binary size

### Release Process

1. Tag release: `git tag v1.0.0`
2. Build for platforms: `GOOS=linux GOARCH=amd64 go build ...`
3. Create GitHub release with binaries

## Code Quality

### Tools

- `gofmt`: Code formatting
- `go vet`: Static analysis
- `golangci-lint`: Comprehensive linting
- `go test`: Unit testing

### CI/CD

GitHub Actions workflow:
1. Test on multiple OS (Linux, macOS, Windows)
2. Test with multiple Go versions
3. Run linters
4. Build binary
5. Security scanning

## Conclusion

The architecture is designed to be:
- **Simple**: Easy to understand and modify
- **Testable**: All components unit tested
- **Extensible**: Easy to add features
- **Maintainable**: Clear separation of concerns
- **Performant**: Fast for typical usage patterns

For questions or suggestions, please open an issue or pull request.
