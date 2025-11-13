#!/bin/bash
# Daily Habit Reminder Script
# Add to crontab: 0 20 * * * /path/to/daily-reminder.sh

# Configuration
HABIT_BIN="${HABIT_BIN:-habit}"
NOTIFY_CMD="${NOTIFY_CMD:-notify-send}"

# Check if any habits were marked today
TODAY=$(date +%Y-%m-%d)
MARKED_TODAY=$($HABIT_BIN list | grep -c "Last done: $TODAY")
TOTAL_HABITS=$($HABIT_BIN list | grep -c "Streak:")

if [ "$MARKED_TODAY" -eq 0 ]; then
    MESSAGE="⚠️  You haven't tracked any habits today! Don't break your streak!"

    # Try to send desktop notification
    if command -v $NOTIFY_CMD &> /dev/null; then
        $NOTIFY_CMD "Habit Tracker" "$MESSAGE"
    fi

    # Print to stdout (useful for cron email)
    echo "$MESSAGE"
    echo ""
    $HABIT_BIN list

elif [ "$MARKED_TODAY" -lt "$TOTAL_HABITS" ]; then
    REMAINING=$((TOTAL_HABITS - MARKED_TODAY))
    MESSAGE="📝 You've tracked $MARKED_TODAY/$TOTAL_HABITS habits today. $REMAINING remaining!"

    if command -v $NOTIFY_CMD &> /dev/null; then
        $NOTIFY_CMD "Habit Tracker" "$MESSAGE"
    fi

    echo "$MESSAGE"
else
    MESSAGE="✅ Great job! All $TOTAL_HABITS habits tracked today!"

    if command -v $NOTIFY_CMD &> /dev/null; then
        $NOTIFY_CMD "Habit Tracker" "$MESSAGE"
    fi

    echo "$MESSAGE"
fi
