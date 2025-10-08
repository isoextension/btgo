package ansi

type Ansi int

const (
	Reset Ansi = iota
	Bold
	Italic
	Red
	Green
	Yellow
	Blue
	White
)

var ansiCodes = map[Ansi]string{
	Reset:  "\x1b[0m",
	Bold:   "\x1b[1m",
	Italic: "\x1b[3m",
	Red:    "\x1b[31m",
	Green:  "\x1b[32m",
	Yellow: "\x1b[33m",
	Blue:   "\x1b[34m",
	White:  "\x1b[37m",
}

func (a Ansi) String() string {
	return ansiCodes[a]
}
