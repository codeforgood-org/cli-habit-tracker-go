# Habit Tracker Examples & Automation Scripts

This directory contains example scripts and configurations for automating your habit tracking workflow.

## Automation Scripts

### Daily Reminder (`daily-reminder.sh`)

Sends you a reminder if you haven't tracked all your habits for the day.

**Setup:**

```bash
# Make executable
chmod +x examples/daily-reminder.sh

# Add to crontab (runs daily at 8 PM)
crontab -e
# Add this line:
0 20 * * * /path/to/cli-habit-tracker-go/examples/daily-reminder.sh
```

**Features:**
- Checks if habits were tracked today
- Sends desktop notification (if `notify-send` is available)
- Provides completion status
- Email output via cron

**Configuration:**
```bash
export HABIT_BIN=/usr/local/bin/habit    # Path to habit binary
export NOTIFY_CMD=notify-send              # Notification command
```

### Weekly Backup (`weekly-backup.sh`)

Automatically backs up your habit data weekly.

**Setup:**

```bash
chmod +x examples/weekly-backup.sh

# Add to crontab (runs every Sunday at midnight)
crontab -e
# Add:
0 0 * * 0 /path/to/cli-habit-tracker-go/examples/weekly-backup.sh
```

**Features:**
- Creates timestamped backups
- Automatic cleanup of old backups (30-day retention by default)
- Configurable backup directory

**Configuration:**
```bash
export BACKUP_DIR=$HOME/habit-backups  # Backup location
export RETENTION_DAYS=30                # Days to keep backups
```

### Weekly Report (`weekly-report.sh`)

Generates a summary report of your habit tracking.

**Setup:**

```bash
chmod +x examples/weekly-report.sh

# Add to crontab (runs every Monday at 9 AM)
crontab -e
# Add:
0 9 * * 1 /path/to/cli-habit-tracker-go/examples/weekly-report.sh
```

**Features:**
- Comprehensive statistics
- Full habit list
- Optional email delivery
- Motivational tips

**Configuration:**
```bash
export REPORT_EMAIL=your@email.com  # Email address for reports
```

## Complete Crontab Example

Here's a complete crontab configuration using all three scripts:

```bash
# Edit crontab
crontab -e

# Add these lines:
# Daily reminder at 8 PM
0 20 * * * /home/user/cli-habit-tracker-go/examples/daily-reminder.sh

# Weekly backup every Sunday at midnight
0 0 * * 0 /home/user/cli-habit-tracker-go/examples/weekly-backup.sh

# Weekly report every Monday at 9 AM
0 9 * * 1 /home/user/cli-habit-tracker-go/examples/weekly-report.sh
```

## Integration Examples

### Shell Aliases

Add these to your `.bashrc` or `.zshrc`:

```bash
# Quick habit aliases
alias ht='habit'
alias htl='habit list'
alias htm='habit mark'
alias hts='habit stats'

# Work vs Personal habits
alias work-habits='HABIT_DATA_FILE=~/work-habits.json habit'
alias personal-habits='HABIT_DATA_FILE=~/personal-habits.json habit'
```

### Git Hooks

Track coding habits automatically:

```bash
# In your project's .git/hooks/post-commit
#!/bin/bash
habit mark "Daily Coding" 2>/dev/null
```

### Tmux Status Bar

Show habit completion in tmux status bar:

```bash
# Add to ~/.tmux.conf
set -g status-right '#(habit stats | grep "Marked today" | cut -d: -f2)'
```

### Desktop Widget

Create a simple desktop widget (requires `conky`):

```bash
# Add to ~/.conkyrc
${execi 300 habit stats}
```

## Advanced Usage

### Multi-Device Sync

Use cloud storage to sync habits across devices:

```bash
# Setup
export HABIT_DATA_FILE=~/Dropbox/habits.json

# Or use git
cd ~/habits-repo
git pull
habit mark "Exercise"
git add habits.json
git commit -m "Update habits"
git push
```

### Daily Habit Checklist

Create a morning checklist script:

```bash
#!/bin/bash
# morning-routine.sh

echo "🌅 Morning Routine Checklist"
echo ""

habits=("Wake up at 6am" "Exercise" "Meditation" "Healthy Breakfast")

for habit in "${habits[@]}"; do
    read -p "Did you: $habit? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        habit mark "$habit"
    fi
done

echo ""
habit stats
```

### Gamification Script

Track points and levels:

```bash
#!/bin/bash
# gamification.sh

TOTAL_STREAK=$(habit stats | grep "Total streak days" | awk '{print $4}')
LEVEL=$((TOTAL_STREAK / 10))
POINTS=$((TOTAL_STREAK * 10))

echo "🎮 Habit Tracker Stats"
echo "Level: $LEVEL"
echo "Points: $POINTS"
echo "Next level in: $((10 - (TOTAL_STREAK % 10))) days"
```

## Troubleshooting

### Cron Not Working

1. Check cron service is running:
   ```bash
   systemctl status cron
   ```

2. Check cron logs:
   ```bash
   grep CRON /var/log/syslog
   ```

3. Ensure scripts have execute permission:
   ```bash
   chmod +x examples/*.sh
   ```

4. Use full paths in cron jobs

### Notifications Not Showing

1. Install notification system:
   ```bash
   # Ubuntu/Debian
   sudo apt-get install libnotify-bin

   # macOS
   brew install terminal-notifier
   ```

2. Set DISPLAY variable in cron:
   ```bash
   DISPLAY=:0
   ```

## Contributing

Have a useful automation script? Please contribute!

1. Add your script to the `examples/` directory
2. Document it in this README
3. Submit a pull request

## License

All example scripts are provided under the same MIT license as the main project.
