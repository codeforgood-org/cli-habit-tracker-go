// Package color provides ANSI color codes for terminal output.
package color

import (
	"fmt"
	"os"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
	White  = "\033[97m"
)

var (
	// NoColor disables color output
	NoColor = false
)

func init() {
	// Disable colors if NO_COLOR environment variable is set
	// or if not running in a terminal
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		NoColor = true
	}
}

// Colorize wraps text in the given color code.
func Colorize(text, colorCode string) string {
	if NoColor {
		return text
	}
	return colorCode + text + Reset
}

// Success returns green colored text.
func Success(text string) string {
	return Colorize(text, Green)
}

// Error returns red colored text.
func Error(text string) string {
	return Colorize(text, Red)
}

// Warning returns yellow colored text.
func Warning(text string) string {
	return Colorize(text, Yellow)
}

// Info returns cyan colored text.
func Info(text string) string {
	return Colorize(text, Cyan)
}

// Highlight returns bold text.
func Highlight(text string) string {
	return Colorize(text, Bold)
}

// Dim returns gray colored text.
func Dim(text string) string {
	return Colorize(text, Gray)
}

// Sprintf formats and colorizes text.
func Sprintf(colorCode, format string, a ...interface{}) string {
	return Colorize(fmt.Sprintf(format, a...), colorCode)
}
