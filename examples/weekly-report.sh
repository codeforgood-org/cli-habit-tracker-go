#!/bin/bash
# Weekly Habit Report Script
# Generates a summary of your habit tracking for the week
# Add to crontab: 0 9 * * 1 /path/to/weekly-report.sh

# Configuration
HABIT_BIN="${HABIT_BIN:-habit}"
EMAIL="${REPORT_EMAIL:-}"

# Generate report
REPORT_FILE="/tmp/habit-report-$(date +%Y%m%d).txt"

{
    echo "==============================================="
    echo "  WEEKLY HABIT TRACKER REPORT"
    echo "  Week of $(date +%Y-%m-%d)"
    echo "==============================================="
    echo ""
    echo "📊 Statistics:"
    echo ""
    $HABIT_BIN stats
    echo ""
    echo "==============================================="
    echo ""
    echo "📋 All Habits:"
    echo ""
    $HABIT_BIN list
    echo ""
    echo "==============================================="
    echo ""
    echo "💡 Tip of the Week:"
    echo "Consistency is key! Even small daily habits"
    echo "compound into significant changes over time."
    echo ""
} > "$REPORT_FILE"

# Display report
cat "$REPORT_FILE"

# Optionally email the report
if [ -n "$EMAIL" ] && command -v mail &> /dev/null; then
    cat "$REPORT_FILE" | mail -s "Weekly Habit Tracker Report" "$EMAIL"
    echo "📧 Report emailed to $EMAIL"
fi

# Cleanup
rm -f "$REPORT_FILE"
