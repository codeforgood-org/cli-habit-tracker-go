#compdef habit

_habit() {
    local -a commands
    commands=(
        'list:List all habits with their streaks'
        'ls:List all habits with their streaks'
        'mark:Mark a habit as done for today'
        'done:Mark a habit as done for today'
        'delete:Delete a habit'
        'del:Delete a habit'
        'rm:Delete a habit'
        'reset:Reset a habit'\''s streak'
        'stats:Show habit statistics'
        'statistics:Show habit statistics'
        'search:Search for habits by name'
        'find:Search for habits by name'
        'edit:Rename a habit'
        'rename:Rename a habit'
        'export:Export habits to a file'
        'import:Import habits from a file'
        'backup:Create a backup of habits data'
        'restore:Restore from a backup file'
        'version:Show version information'
        'help:Show detailed help'
    )

    local -a export_formats
    export_formats=(
        'csv:Export as CSV format'
        'json:Export as JSON format'
    )

    local -a import_formats
    import_formats=(
        'csv:Import from CSV format'
        'json:Import from JSON format'
    )

    _arguments -C \
        '1: :->command' \
        '*::arg:->args'

    case "$state" in
        command)
            _describe 'command' commands
            ;;
        args)
            case "${words[1]}" in
                export)
                    if [[ $CURRENT -eq 2 ]]; then
                        _describe 'format' export_formats
                    elif [[ $CURRENT -eq 3 ]]; then
                        _files
                    fi
                    ;;
                import)
                    if [[ $CURRENT -eq 2 ]]; then
                        _describe 'format' import_formats
                    elif [[ $CURRENT -eq 3 ]]; then
                        _files
                    elif [[ $CURRENT -eq 4 ]]; then
                        _values 'options' '--merge[Merge with existing habits]'
                    fi
                    ;;
                backup|restore)
                    _files
                    ;;
            esac
            ;;
    esac
}

_habit "$@"
