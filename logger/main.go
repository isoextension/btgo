package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/isoextension/btgo/ansi"
)

type Logger struct {
	prefix string
}

func New(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

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

func (l *Logger) Fplainf(stream io.Writer, ptrn string, objs ...any) {
	fmt.Fprintf(
		stream,
		ptrn,
		objs...,
	)
}

func (l *Logger) Fplain(stream io.Writer, objs ...any) {
	fmt.Fprint(
		stream,
		objs...,
	)
}

func (l *Logger) Plain(objs ...any) {
	fmt.Print(objs...)
}

func (l *Logger) Plainln(objs ...any) {
	fmt.Println(objs...)
}

func (l *Logger) Info(objs ...any) {
	l.Logln(os.Stdout, "INFO", ansi.White, str)
}

func (l *Logger) Debug(objs ...any) {
	l.Logln(os.Stdout, "DEBUG", ansi.Blue, objs...)
}

func (l *Logger) Warning(objs ...any) {
	l.Logln(os.Stderr, "WARN", ansi.Yellow, objs...)
}

func (l *Logger) Error(objs ...any) {
	l.Logln(os.Stderr, "ERR", ansi.Red, objs...)
}

func (l *Logger) Major(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, "  %s%s==>%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

func (l *Logger) Minor(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, " %s%s->%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}

func (l *Logger) Colon(color ansi.Ansi, objs ...any) {
	l.Fplainf(os.Stdout, " %s%s::%s %s\n", ansi.Bold, color, ansi.Reset, fmt.Sprint(objs...))
}
