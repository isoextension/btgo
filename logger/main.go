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

func (l *Logger) Info(str string) {
	l.Logln(os.Stdout, "INFO", ansi.White, str)
}

func (l *Logger) Debug(str string) {
	l.Logln(os.Stdout, "DEBUG", ansi.Blue, str)
}

func (l *Logger) Warning(str string) {
	l.Logln(os.Stderr, "WARN", ansi.Yellow, str)
}

func (l *Logger) Error(str string) {
	l.Logln(os.Stderr, "ERR", ansi.Red, str)
}

func (l *Logger) Major(str string, color ansi.Ansi) {
	l.Fplainf(os.Stdout, "  %s%s==>%s %s\n", ansi.Bold, color, ansi.Reset, str)
}

func (l *Logger) Minor(str string, color ansi.Ansi) {
	l.Fplainf(os.Stdout, " %s%s->%s %s\n", ansi.Bold, color, ansi.Reset, str)
}

func (l *Logger) Colon(str string, color ansi.Ansi) {
	l.Fplainf(os.Stdout, " %s%s::%s %s\n", ansi.Bold, color, ansi.Reset, str)
}
