# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-01-13

### Added

**Major Features:**
- **Export command**: Export habits to CSV or JSON formats
- **Import command**: Import habits from CSV or JSON files with merge support
- **Search command**: Search for habits by name (case-insensitive substring match)
- **Edit command**: Rename habits without losing streak data
- **Backup command**: Create timestamped backups of habit data
- **Restore command**: Restore habits from backup files with automatic current data backup

**Developer Experience:**
- Shell completion scripts for Bash, Zsh, and Fish
- GoReleaser configuration for automated multi-platform releases
- Dockerfile for Docker support
- Installation script for quick setup
- Comprehensive FAQ documentation (docs/FAQ.md)
- Enhanced README with badges, advanced examples, and full command reference

**Project Infrastructure:**
- Added .goreleaser.yml for release automation
- Added Dockerfile for containerized usage
- Added install.sh for one-command installation
- Added shell completion scripts in completions/
- Added comprehensive FAQ in docs/FAQ.md

### Changed

**Breaking Changes:**
- Version bumped to 2.0.0 to reflect major feature additions
- Data file default location changed to `~/.habit-tracker/habits.json` (was `./habits.json`)

**Improvements:**
- Enhanced help command with detailed documentation for all commands
- Improved usage output to categorize core and advanced commands
- Better error messages across all commands
- Updated README with comprehensive documentation and examples

**Refactoring:**
- All new commands follow consistent error handling patterns
- Improved code organization with dedicated command files

### Fixed

- Fixed unused variable in backup.go restore function

## [1.0.0] - 2025-01-13

### Added

**Initial Release:**
- Core habit tracking functionality
- List command to view all habits
- Mark command to complete habits daily
- Delete command to remove habits
- Reset command to clear habit streaks
- Stats command to view habit statistics
- Automatic streak calculation for consecutive days
- JSON-based persistent storage
- Configurable data file location via environment variable
- Comprehensive unit tests for models and storage

**Project Structure:**
- Modular architecture with clear separation of concerns:
  - `cmd/habit`: Main application entry point
  - `pkg/models`: Data models and business logic
  - `pkg/storage`: JSON persistence layer
  - `pkg/commands`: CLI command handlers
  - `internal/config`: Configuration management

**Documentation:**
- README with installation and usage instructions
- CONTRIBUTING.md with development guidelines
- ARCHITECTURE.md with design documentation
- LICENSE (MIT)
- Inline code documentation

**Build & Development:**
- Makefile for common tasks (build, test, lint, install)
- .gitignore for Go projects
- .golangci.yml for code quality checks
- GitHub Actions CI/CD workflow for testing, linting, and building
- Multi-platform testing (Linux, macOS, Windows)
- Multiple Go version testing (1.21, 1.22)

### Features

- Track unlimited habits with daily completion
- Automatic streak calculation
- Case-insensitive habit name matching
- Persistent JSON storage
- Configurable data file location
- Comprehensive statistics (total, max streak, average streak, completion rate)
- Cross-platform support (Linux, macOS, Windows)

---

## Release Links

- [2.0.0](https://github.com/codeforgood-org/cli-habit-tracker-go/releases/tag/v2.0.0) - 2025-01-13
- [1.0.0](https://github.com/codeforgood-org/cli-habit-tracker-go/releases/tag/v1.0.0) - 2025-01-13

## Upgrade Guide

### Upgrading from 1.0.0 to 2.0.0

**Data File Location Change:**

The default data file location has changed from `./habits.json` to `~/.habit-tracker/habits.json`.

To migrate your existing data:

```bash
# If you have habits.json in your current directory
mkdir -p ~/.habit-tracker
mv ./habits.json ~/.habit-tracker/habits.json
```

Or, continue using the old location by setting an environment variable:

```bash
export HABIT_DATA_FILE=./habits.json
```

**New Commands:**

All existing commands continue to work as before. New commands have been added:
- `search`, `edit`, `export`, `import`, `backup`, `restore`

No changes required to your existing workflows!

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute to this project.
