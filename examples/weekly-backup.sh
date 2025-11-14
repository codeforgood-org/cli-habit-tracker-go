#!/bin/bash
# Weekly Habit Backup Script
# Add to crontab: 0 0 * * 0 /path/to/weekly-backup.sh

# Configuration
HABIT_BIN="${HABIT_BIN:-habit}"
BACKUP_DIR="${BACKUP_DIR:-$HOME/habit-backups}"
RETENTION_DAYS="${RETENTION_DAYS:-30}"

# Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# Create timestamped backup
TIMESTAMP=$(date +%Y%m%d-%H%M%S)
BACKUP_FILE="$BACKUP_DIR/habits-weekly-$TIMESTAMP.json"

# Run backup
echo "Creating weekly backup..."
$HABIT_BIN backup "$BACKUP_FILE"

if [ $? -eq 0 ]; then
    echo "✅ Backup created successfully: $BACKUP_FILE"

    # Delete old backups (older than RETENTION_DAYS)
    echo "Cleaning up old backups (older than $RETENTION_DAYS days)..."
    find "$BACKUP_DIR" -name "habits-weekly-*.json" -type f -mtime +$RETENTION_DAYS -delete

    # Count remaining backups
    BACKUP_COUNT=$(find "$BACKUP_DIR" -name "habits-weekly-*.json" | wc -l)
    echo "Total backups: $BACKUP_COUNT"
else
    echo "❌ Backup failed!"
    exit 1
fi
