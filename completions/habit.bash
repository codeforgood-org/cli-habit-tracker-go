#!/usr/bin/env bash
# Bash completion script for habit tracker

_habit_completions() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    # All available commands
    local commands="list ls mark done delete del rm reset stats statistics search find edit rename export import backup restore version help"

    # Command-specific completions
    case "${prev}" in
        habit)
            COMPREPLY=( $(compgen -W "${commands}" -- ${cur}) )
            return 0
            ;;
        export)
            COMPREPLY=( $(compgen -W "csv json" -- ${cur}) )
            return 0
            ;;
        import)
            COMPREPLY=( $(compgen -W "csv json" -- ${cur}) )
            return 0
            ;;
        csv|json)
            # Offer file completion
            COMPREPLY=( $(compgen -f -- ${cur}) )
            return 0
            ;;
        backup|restore)
            # Offer file completion
            COMPREPLY=( $(compgen -f -- ${cur}) )
            return 0
            ;;
        *)
            ;;
    esac

    # Default to command completion
    if [ $COMP_CWORD -eq 1 ]; then
        COMPREPLY=( $(compgen -W "${commands}" -- ${cur}) )
    fi
}

complete -F _habit_completions habit
