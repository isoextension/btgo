// Package taylog provides a simple structured logging system
// with wrappers for fmt, os, and custom formatting options.
package taylog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/isoextension/btgo/ansi"
)

// Logger is the main logger instance, used for structured and formatted logging.
type Logger struct {
	prefix string
	mutex  sync.Mutex
	stream *io.Writer
}

// Level is the log level... of course
type Level int

const (
	Trace Level = iota // 0
	Debug              // 1
	Info               // 2
	Warning            // 3
	Error              // 4
	Fatal              // 5
)

// New creates and returns a new Logger with the given prefix.
func New(prefix string, stream *io.Writer) *Logger {
	return &Logger{prefix: prefix, stream: stream}
}

/// ESSENTIAL ///

// Logln writes a formatted log line to the specified stream,
// including a log level, color, and automatic newline.
func (l *Logger) Logln(stream io.Writer, level string, color ansi.Ansi, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Fprintf(
		stream,
		"[%s%s%s] [%s]: %s\n",
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
	l.mutex.Lock()
	defer l.mutex.Unlock()
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

/// PLAIN ///

// Fplainf writes formatted text to a specified stream without
// additional log decorations or automatic newlines.
func (l *Logger) Fplainf(stream io.Writer, ptrn string, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Fprintf(stream, ptrn, objs...)
}

// Fplain writes plain text to a specified stream without additional log decorations
// or automatic newlines.
func (l *Logger) Fplain(stream io.Writer, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Fprint(stream, objs...)
}

// Plain writes plain text directly to standard output without formatting.
func (l *Logger) Plain(objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Print(objs...)
}

func (l *Logger) Plainf(pattern string, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Printf(pattern, objs...)
}

// Plainln writes plain text directly to standard output,
// automatically adding a newline.
func (l *Logger) Plainln(objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	fmt.Println(objs...)
}

/// LEVELED ///

// Info logs informational messages to stdout
func (l *Logger) Info(objs ...any) {
    l.Logln(os.Stdout, "INFO", ansi.White, objs...)
}

// Infof logs informational messages to stdout with a format string
func (l *Logger) Infof(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Logln(os.Stdout, "INFO", ansi.White, []any{message})
}

// Debug logs verbose debug messages to stdout.
// Typically used with a --debug flag or development builds.
func (l *Logger) Debug(objs ...any) {
    l.Logln(os.Stdout, "DEBUG", ansi.Blue, objs...)
}

// Debugf logs verbose debug messages to stdout with a format string.
// Typically used with a --debug flag or development builds.
func (l *Logger) Debugf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Logln(os.Stdout, "DEBUG", ansi.Blue, []any{message})
}

// Warning logs warning messages to stderr.
// Used for non-critical issues or potential problems.
func (l *Logger) Warning(objs ...any) {
    l.Logln(os.Stderr, "WARN", ansi.Yellow, objs...)
}

// Warningf logs warning messages to stderr with a format string.
// Used for non-critical issues or potential problems.
func (l *Logger) Warningf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Logln(os.Stderr, "WARN", ansi.Yellow, []any{message})
}

// Error logs error messages to stderr.
// Used for serious issues or failures.
func (l *Logger) Error(objs ...any) {
    l.Logln(os.Stderr, "ERROR", ansi.Red, objs...)
}

// Errorf logs error messages to stderr with a format string.
// Used for serious issues or failures.
func (l *Logger) Errorf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Logln(os.Stderr, "ERROR", ansi.Red, []any{message})
}

/// BASIC ///

// BasicInfo prints like Info with a different format.  i  hello world
func (l *Logger) BasicInfo(objs ...any) {
    l.Fplainf(os.Stdout, "%si%s  %s\n", ansi.BrightWhite, ansi.Reset, fmt.Sprint(objs...))
}

// BasicInfof prints like Info with a different format and a format string.  i  hello world
func (l *Logger) BasicInfof(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Fplainf(os.Stdout, "%si%s  %s\n", ansi.BrightWhite, ansi.Reset, message)
}

// BasicWarn prints like Warn with a different format   !  something is suspicious
func (l *Logger) BasicWarn(objs ...any) {
    l.Fplainf(os.Stderr, "%s!%s %s\n", ansi.BrightYellow.String(), ansi.Reset, fmt.Sprint(objs...))
}

// BasicWarnf prints like Warn with a different format and a format string   !  something is suspicious
func (l *Logger) BasicWarnf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Fplainf(os.Stderr, "%s!%s %s\n", ansi.BrightYellow.String(), ansi.Reset, message)
}

// BasicError prints like Error with a different format   x  oopsies
func (l *Logger) BasicError(objs ...any) {
    l.Fplainf(os.Stdout, "%sx%s %s\n", ansi.BrightRed.String(), ansi.Reset, fmt.Sprint(objs...))
}

// BasicErrorf prints like Error with a different format and a format string   x  oopsies
func (l *Logger) BasicErrorf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Fplainf(os.Stdout, "%sx%s %s", ansi.BrightRed.String(), ansi.Reset, message)
}

// BasicFatal prints like Fatal with a different format   x_x  oopsies
func (l *Logger) BasicFatal(objs ...any) {
    l.Fplainf(os.Stdout, "%s%sx_x%s %s\n", ansi.Bold.String(), ansi.Red.String(), ansi.Reset, fmt.Sprint(objs...))
}

// BasicFatalf prints like Fatal with a different format and a format string   x_x  oopsies
func (l *Logger) BasicFatalf(format string, objs ...any) {
    message := fmt.Sprintf(format, objs...)
    l.Fplainf(os.Stdout, "%s%sx_x%s %s", ansi.Bold.String(), ansi.Red.String(), ansi.Reset, message)
}

/// SYSTEM STYLED ///

// Major logs major events such as warnings, errors, or important tasks,
// styled in mkinitcpio-like formatting.
func (l *Logger) Major(color ansi.Ansi, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.Fplainf(os.Stdout, "  %s%s=>%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

// Minor logs minor events such as informational or debug messages,
// styled in mkinitcpio-like formatting.
func (l *Logger) Minor(color ansi.Ansi, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.Fplainf(os.Stdout, " %s%s->%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

// Colon logs messages using pacman-style "::" formatting.
func (l *Logger) Colon(color ansi.Ansi, objs ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.Fplainf(os.Stdout, " %s%s::%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

/// MISC ///

func (l *Logger) SetPrefix(prefix string) (*Logger, error) {
	if prefix == "" {
		return l, errors.New("prefix cannot be empty")
	}
	l.prefix = prefix
	return l, nil
}

func (l *Logger) SetDefaultStream(stream *io.Writer) *Logger {
	l.stream = stream
	return l
}
