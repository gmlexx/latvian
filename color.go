package main

import "os"

// Minimal ANSI color helpers. Colors are disabled automatically if the
// NO_COLOR environment variable is set (see https://no-color.org/) or if
// stdout does not look like a terminal isn't checked (kept simple on
// purpose) — most modern terminals on Linux/macOS/WSL support ANSI codes.
const (
	ansiReset  = "\033[0m"
	ansiBold   = "\033[1m"
	ansiRed    = "\033[31m"
	ansiGreen  = "\033[32m"
	ansiYellow = "\033[33m"
	ansiCyan   = "\033[36m"
)

var colorsEnabled = os.Getenv("NO_COLOR") == ""

func colorize(code, s string) string {
	if !colorsEnabled {
		return s
	}
	return code + s + ansiReset
}

func green(s string) string  { return colorize(ansiGreen, s) }
func red(s string) string    { return colorize(ansiRed, s) }
func yellow(s string) string { return colorize(ansiYellow, s) }
func cyan(s string) string   { return colorize(ansiCyan, s) }
func bold(s string) string   { return colorize(ansiBold, s) }
