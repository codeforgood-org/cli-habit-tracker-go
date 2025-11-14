# Frequently Asked Questions (FAQ)

## General Questions

### What is Habit Tracker?

Habit Tracker is a simple command-line tool for tracking daily habits and building streaks. It helps you maintain consistency by tracking consecutive days of completing habits.

### Is it free?

Yes! Habit Tracker is open source and completely free to use under the MIT License.

### What platforms are supported?

Habit Tracker supports:
- Linux (amd64, arm64, armv7)
- macOS (amd64, arm64/M1/M2)
- Windows (amd64)

## Installation

### How do I install Habit Tracker?

**Using the install script (recommended):**
```bash
curl -sfL https://raw.githubusercontent.com/codeforgood-org/cli-habit-tracker-go/main/install.sh | sh
```

**Using Go:**
```bash
go install github.com/codeforgood-org/cli-habit-tracker-go/cmd/habit@latest
```

**Manual installation:**
Download the binary for your platform from the [releases page](https://github.com/codeforgood-org/cli-habit-tracker-go/releases).

### Where is the data stored?

By default, habits are stored in `~/.habit-tracker/habits.json`. You can customize this location using the `HABIT_DATA_FILE` environment variable:

```bash
export HABIT_DATA_FILE=~/my-habits.json
```

### How do I enable shell completions?

**Bash:**
```bash
cp completions/habit.bash /etc/bash_completion.d/habit
# or
cp completions/habit.bash ~/.bash_completion.d/habit
```

**Zsh:**
```bash
cp completions/habit.zsh ~/.zsh/completions/_habit
# Add to ~/.zshrc:
fpath=(~/.zsh/completions $fpath)
```

**Fish:**
```bash
cp completions/habit.fish ~/.config/fish/completions/
```

## Usage Questions

### How do I create a new habit?

Just mark it as done! If the habit doesn't exist, it will be created automatically:

```bash
habit mark "Morning Exercise"
```

### How does streak calculation work?

- **Day 1**: Mark a habit for the first time → Streak = 1
- **Day 2**: Mark it the next day (consecutive) → Streak = 2
- **Day 3**: Mark it the day after → Streak = 3
- **Gap**: Skip a day → Streak resets to 1 when you mark it again

Streaks only increment on consecutive days (based on dates, not 24-hour periods).

### Can I track multiple habits?

Yes! You can track as many habits as you want. Each habit has its own independent streak.

### Can habit names have spaces?

Yes! Habit names can contain spaces, emojis, and most special characters:

```bash
habit mark "30 minutes reading 📚"
habit mark "Drink 8 glasses of water"
```

### How do I delete a habit?

```bash
habit delete "Habit Name"
# or
habit rm "Habit Name"
```

### Can I rename a habit?

Yes, use the edit command:

```bash
habit edit "Old Name" "New Name"
```

### How do I reset a habit's streak without deleting it?

```bash
habit reset "Habit Name"
```

This sets the streak to 0 and clears the last done date, but keeps the habit in your list.

## Data Management

### How do I backup my habits?

```bash
# Auto-generated filename with timestamp
habit backup

# Custom filename
habit backup my-habits-backup.json
```

### How do I restore from a backup?

```bash
habit restore my-habits-backup.json
```

Your current data is automatically backed up before restoring.

### Can I export my habits to CSV for Excel?

Yes!

```bash
habit export csv habits.csv
```

### Can I import habits from a CSV file?

Yes!

```bash
# Replace existing habits
habit import csv habits.csv

# Merge with existing habits
habit import csv habits.csv --merge
```

### How do I sync habits across multiple devices?

Currently, there's no built-in sync. However, you can:

1. **Manual sync**: Use export/import or backup/restore
2. **Cloud storage**: Set `HABIT_DATA_FILE` to a cloud-synced folder:
   ```bash
   export HABIT_DATA_FILE=~/Dropbox/habits.json
   ```
3. **Git**: Store your habits.json in a git repository

## Troubleshooting

### "Permission denied" when installing

The install script needs write access to `/usr/local/bin`. Either:

1. Run with sudo: `curl -sfL ... | sudo sh`
2. Install to a user directory:
   ```bash
   mkdir -p ~/bin
   # Download and extract manually to ~/bin
   # Add ~/bin to PATH
   ```

### "command not found: habit" after installation

The installation directory may not be in your PATH. Add it to your shell's RC file:

```bash
# For bash (~/.bashrc) or zsh (~/.zshrc)
export PATH="$PATH:/usr/local/bin"

# Reload your shell
source ~/.bashrc  # or source ~/.zshrc
```

### My habit streak reset unexpectedly

This happens when there's a gap of more than one day. Streaks only increment on consecutive days.

**Example:**
- Day 1 (Monday): Mark habit → Streak = 1
- Day 2 (Tuesday): Mark habit → Streak = 2
- Day 3 (Wednesday): **Skip**
- Day 4 (Thursday): Mark habit → Streak resets to 1

### I marked a habit twice on the same day

That's okay! The second mark is ignored with a message "Already marked today."

### Can I mark a habit for a past date?

No, currently the tool only supports marking habits for the current date. This is intentional to prevent backdating and maintain integrity.

### The data file is corrupted

If your `habits.json` file gets corrupted:

1. Check if you have a backup (auto-generated during restore operations)
2. Try manually fixing the JSON syntax
3. If all else fails, start fresh (the file will be recreated automatically)

### How do I completely uninstall Habit Tracker?

```bash
# Remove the binary
sudo rm /usr/local/bin/habit

# Remove data (optional)
rm -rf ~/.habit-tracker

# Remove completions (optional)
rm /etc/bash_completion.d/habit
rm ~/.zsh/completions/_habit
rm ~/.config/fish/completions/habit.fish
```

## Advanced Usage

### Can I use this in scripts?

Yes! All commands return appropriate exit codes:
- 0: Success
- 1: Error

Example script:
```bash
#!/bin/bash
habit mark "Daily Script Run" && echo "Logged successfully"
```

### Can I run multiple instances?

Yes, as long as they use different data files:

```bash
HABIT_DATA_FILE=~/work-habits.json habit mark "Code Review"
HABIT_DATA_FILE=~/personal-habits.json habit mark "Exercise"
```

### How do I view habits in JSON format programmatically?

```bash
cat ~/.habit-tracker/habits.json | jq .
```

Or export to JSON:
```bash
habit export json habits.json
cat habits.json | jq .
```

### Can I add custom fields to habits?

Not through the CLI currently. However, the JSON format is extensible. You can manually edit `habits.json` to add custom fields (they'll be preserved but not used by the CLI).

### How do I generate reports?

Use the stats command:
```bash
habit stats
```

For custom reports, export to CSV and use Excel, Google Sheets, or command-line tools:
```bash
habit export csv habits.csv
# Then open in Excel or use csvkit, pandas, etc.
```

## Feature Requests

### Will you add feature X?

Check our [GitHub Issues](https://github.com/codeforgood-org/cli-habit-tracker-go/issues) to see if it's already requested. If not, please create a new issue!

Popular requested features:
- [ ] Cloud sync
- [ ] Reminders/notifications
- [ ] Charts and visualizations
- [ ] Goals (e.g., "30-day streak")
- [ ] Habit categories/tags
- [ ] Weekly/monthly views
- [ ] Mobile app

### How can I contribute?

See our [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines. We welcome:
- Bug reports
- Feature requests
- Code contributions
- Documentation improvements
- Translation

## Getting Help

### Where can I get help?

1. Read this FAQ
2. Check the [documentation](../README.md)
3. Search [existing issues](https://github.com/codeforgood-org/cli-habit-tracker-go/issues)
4. Create a [new issue](https://github.com/codeforgood-org/cli-habit-tracker-go/issues/new)

### How do I report a bug?

Create an issue on GitHub with:
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version, habit version)
- Error messages or logs

## Privacy & Security

### Is my data private?

Yes! All data is stored locally on your computer. Nothing is sent to external servers unless you explicitly sync to a cloud service.

### Is the data encrypted?

Not by default. The `habits.json` file is plain text. If you need encryption:
1. Use encrypted disk/partition
2. Use encrypted cloud storage
3. Manually encrypt the file

### Can others see my habits?

Only if they have access to your computer or the file location you've specified. Make sure to set appropriate file permissions:

```bash
chmod 600 ~/.habit-tracker/habits.json
```

---

**Didn't find your answer?** Create an issue on [GitHub](https://github.com/codeforgood-org/cli-habit-tracker-go/issues)!
