// Package logger provides a simple structured logging system
// with wrappers for fmt, os, and custom formatting options.
package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/isoextension/btgo/ansi"
)

// Logger is the main logger instance, used for structured and formatted logging.
type Logger struct {
	prefix string
}

// New creates and returns a new Logger with the given prefix.
func New(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Logln writes a formatted log line to the specified stream,
// including a log level, color, and automatic newline.
func (l *Logger) Logln(stream io.Writer, level string, color ansi.Ansi, objs ...any) {
	fmt.Fprintf(
		stream,
		"[%s%s%s%s] [%s]: %s\n",
		ansi.Bold.String(),
		color.String(),
		level,
		ansi.Reset.String(),
		l.prefix,
		fmt.Sprint(objs...),
	)
}

// Logf writes a formatted log line to the specified stream,
// similar to fmt.Fprintf, but with a prefixed and colored log header.
func (l *Logger) Logf(stream io.Writer, level string, color ansi.Ansi, ptrn string, objs ...any) {
	fmt.Fprintf(
		stream,
		"[%s%s%s%s] [%s]: %s",
		ansi.Bold.String(),
		color.String(),
		level,
		ansi.Reset.String(),
		l.prefix,
		fmt.Sprintf(ptrn, objs...),
	)
}

// Fplainf writes formatted text to a specified stream without
// additional log decorations or automatic newlines.
func (l *Logger) Fplainf(stream io.Writer, ptrn string, objs ...any) {
	fmt.Fprintf(stream, ptrn, objs...)
}

// Fplain writes plain text to a specified stream without formatting
// or automatic newlines.
func (l *Logger) Fplain(stream io.Writer, objs ...any) {
	fmt.Fprint(stream, objs...)
}

// Plain writes plain text directly to standard output without formatting.
func (l *Logger) Plain(objs ...any) {
	fmt.Print(objs...)
}

// Plainln writes plain text directly to standard output,
// automatically adding a newline.
func (l *Logger) Plainln(objs ...any) {
	fmt.Println(objs...)
}

// Info logs informational messages to standard output.
func (l *Logger) Info(objs ...any) {
	l.Logln(os.Stdout, "INFO", ansi.White, objs...)
}

// Debug logs verbose debug messages to standard output.
// Typically used with a --debug flag or development builds.
func (l *Logger) Debug(objs ...any) {
	l.Logln(os.Stdout, "DEBUG", ansi.Blue, objs...)
}

// Warning logs warning messages to standard error.
// Used for non-critical issues or potential problems.
func (l *Logger) Warning(objs ...any) {
	l.Logln(os.Stderr, "WARN", ansi.Yellow, objs...)
}

// Error logs error messages to standard error.
// Used for serious issues or failures.
func (l *Logger) Error(objs ...any) {
	l.Logln(os.Stderr, "ERR", ansi.Red, objs...)
}

// Major logs major events such as warnings, errors, or important tasks,
// styled in mkinitcpio-like formatting.
func (l *Logger) Major(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, "  %s%s=>%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

// Minor logs minor events such as informational or debug messages,
// styled in mkinitcpio-like formatting.
func (l *Logger) Minor(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, " %s%s->%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

// Colon logs messages using pacman-style "::" formatting.
func (l *Logger) Colon(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, " %s%s::%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}
