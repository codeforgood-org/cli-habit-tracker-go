# Fish completion for habit tracker

# Core commands
complete -c habit -f -n "__fish_use_subcommand" -a "list" -d "List all habits with their streaks"
complete -c habit -f -n "__fish_use_subcommand" -a "ls" -d "List all habits with their streaks"
complete -c habit -f -n "__fish_use_subcommand" -a "mark" -d "Mark a habit as done for today"
complete -c habit -f -n "__fish_use_subcommand" -a "done" -d "Mark a habit as done for today"
complete -c habit -f -n "__fish_use_subcommand" -a "delete" -d "Delete a habit"
complete -c habit -f -n "__fish_use_subcommand" -a "del" -d "Delete a habit"
complete -c habit -f -n "__fish_use_subcommand" -a "rm" -d "Delete a habit"
complete -c habit -f -n "__fish_use_subcommand" -a "reset" -d "Reset a habit's streak"
complete -c habit -f -n "__fish_use_subcommand" -a "stats" -d "Show habit statistics"
complete -c habit -f -n "__fish_use_subcommand" -a "statistics" -d "Show habit statistics"

# Advanced commands
complete -c habit -f -n "__fish_use_subcommand" -a "search" -d "Search for habits by name"
complete -c habit -f -n "__fish_use_subcommand" -a "find" -d "Search for habits by name"
complete -c habit -f -n "__fish_use_subcommand" -a "edit" -d "Rename a habit"
complete -c habit -f -n "__fish_use_subcommand" -a "rename" -d "Rename a habit"
complete -c habit -f -n "__fish_use_subcommand" -a "export" -d "Export habits to a file"
complete -c habit -f -n "__fish_use_subcommand" -a "import" -d "Import habits from a file"
complete -c habit -f -n "__fish_use_subcommand" -a "backup" -d "Create a backup of habits data"
complete -c habit -f -n "__fish_use_subcommand" -a "restore" -d "Restore from a backup file"

# Other commands
complete -c habit -f -n "__fish_use_subcommand" -a "version" -d "Show version information"
complete -c habit -f -n "__fish_use_subcommand" -a "help" -d "Show detailed help"

# Export format completions
complete -c habit -f -n "__fish_seen_subcommand_from export" -a "csv" -d "Export as CSV"
complete -c habit -f -n "__fish_seen_subcommand_from export" -a "json" -d "Export as JSON"

# Import format completions
complete -c habit -f -n "__fish_seen_subcommand_from import" -a "csv" -d "Import from CSV"
complete -c habit -f -n "__fish_seen_subcommand_from import" -a "json" -d "Import from JSON"

# Import merge flag
complete -c habit -f -n "__fish_seen_subcommand_from import" -l merge -d "Merge with existing habits"
