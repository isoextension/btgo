// Package ansi is a library for ansi color/control codes
package ansi

import (
	"fmt"
	"math"
	"strings"
)

// Ansi is a type for referencing an ansi color/control code
type Ansi int

const (
	Red Ansi = iota + 30
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	BgBlack Ansi = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

const (
	BrightBlack Ansi = iota + 90
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

const (
	BgBrightBlack Ansi = iota + 100
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWhite
)

const (
	Reset Ansi = iota
	Bold
	Dim
	Italic
	Underline
	SlowBlink
	FastBlink
	Invert
	Hidden
	Strikethrough
)

const (
	IntensityReset = iota + 20
	ItalicReset
	UnderlineReset
	BlinkReset
	InvertReset
	Show
	StrikeReset
)

var Codes = map[Ansi]string{
	// colors
	Reset:             "\x1b[0m",
	Bold:              "\x1b[1m",
	Dim:               "\x1b[2m",
	Italic:            "\x1b[3m",
	Underline:         "\x1b[4m",
	SlowBlink:         "\x1b[5m",
	FastBlink:         "\x1b[6m",
	Invert:            "\x1b[7m",
	Hidden:            "\x1b[8m",
	Strikethrough:     "\x1b[9m",


	// resets
	IntensityReset:    "\x1b[22m",
	ItalicReset:       "\x1b[23m",
	UnderlineReset:    "\x1b[24m",
	BlinkReset:        "\x1b[25m",
	InvertReset:       "\x1b[27m",
	Show:              "\x1b[28m",
	StrikeReset:       "\x1b[29m",

	// 8/16 colors
	Red:               "\x1b[31m",
	Green:             "\x1b[32m",
	Yellow:            "\x1b[33m",
	Blue:              "\x1b[34m",
	Magenta:           "\x1b[35m",
	Cyan:              "\x1b[36m",
	White:             "\x1b[37m",
	BgBlack:           "\x1b[40m",
	BgRed:             "\x1b[41m",
	BgGreen:           "\x1b[42m",
	BgYellow:          "\x1b[43m",
	BgBlue:            "\x1b[44m",
	BgMagenta:         "\x1b[45m",
	BgCyan:            "\x1b[46m",
	BgWhite:           "\x1b[47m",

	// bright colors
	BrightBlack:       "\x1b[90m",
	BrightRed:         "\x1b[91m",
	BrightGreen:       "\x1b[92m",
	BrightYellow:      "\x1b[93m",
	BrightBlue:        "\x1b[94m",
	BrightMagenta:     "\x1b[95m",
	BrightCyan:        "\x1b[96m",
	BrightWhite:       "\x1b[97m",
	BgBrightBlack:     "\x1b[100m",
	BgBrightRed:       "\x1b[101m",
	BgBrightGreen:     "\x1b[102m",
	BgBrightYellow:    "\x1b[103m",
	BgBrightBlue:      "\x1b[104m",
	BgBrightMagenta:   "\x1b[105m",
	BgBrightCyan:      "\x1b[106m",
	BgBrightWhite:     "\x1b[107m",
}

// String returns the ansi code of a
func (a Ansi) String() string {
	return Codes[a]
}

// Wrap wraps strs... in a color
func (a Ansi) Wrap(strs ...string) string {
	return a.String() + strings.Join(strs, " ") + Codes[Reset]
}

// Brighten returns the bright variant of color
func Brighten(color Ansi) string {
	return color.Bright()
}

// Bright returns the bright variant of a
func (a Ansi) Bright() string {
	// Map regular colors to their bright equivalents
	switch a {
	case Red:
		return Codes[BrightRed]
	case Green:
		return Codes[BrightGreen]
	case Yellow:
		return Codes[BrightYellow]
	case Blue:
		return Codes[BrightBlue]
	case Magenta:
		return Codes[BrightMagenta]
	case Cyan:
		return Codes[BrightCyan]
	case White:
		return Codes[BrightWhite]
	case BgRed:
		return Codes[BgBrightRed]
	case BgGreen:
		return Codes[BgBrightGreen]
	case BgYellow:
		return Codes[BgBrightYellow]
	case BgBlue:
		return Codes[BgBrightBlue]
	case BgMagenta:
		return Codes[BgBrightMagenta]
	case BgCyan:
		return Codes[BgBrightCyan]
	case BgWhite:
		return Codes[BgBrightWhite]
	default:
		// For colors that are already bright or don't have bright variants
		return Codes[a]
	}
}

func Rgb(r, g, b uint8, bg bool) string {
	var prefix int
	prefix = 38
	if bg {
		prefix = 48
	}
	return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", prefix, r, g, b)
}

func OkhslToRgb(h, s, l float64) (float64, float64, float64) {
	// Convert OKHSL to OKLab
	L := l
	a := s * math.Cos(h*math.Pi/180) // Convert hue to Cartesian coordinate 'a'
	b := s * math.Sin(h*math.Pi/180) // Convert hue to Cartesian coordinate 'b'

	// Convert OKLab to linear RGB
	// These coefficients are derived from the OKLab color space transformation matrix
	linearL := L + 0.3963377774*a + 0.2158037573*b // Calculate linear L component
	linearM := L - 0.1055613458*a - 0.0638541728*b // Calculate linear M component
	linearS := L - 0.0894841775*a - 1.2914855480*b // Calculate linear S component

	// Apply the inverse matrix to convert from linear RGB to XYZ
	// These coefficients are derived from the XYZ to linear RGB transformation matrix
	linearL = linearL * 1.9106812443 // Scale linear L component
	linearM = linearM * 1.0000000000 // Scale linear M component
	linearS = linearS * 1.5512010052 // Scale linear S component

	// Convert XYZ to linear RGB
	// These coefficients are derived from the XYZ to linear RGB transformation matrix
	r := 3.2404542*linearL - 1.5371385*linearM - 0.4985314*linearS // Calculate red component
	g := -0.9692660*linearL + 1.8760108*linearM + 0.0415560*linearS // Calculate green component
	linearB := 0.0556434*linearL - 0.2040259*linearM + 1.0572252*linearS // Calculate blue component

	// Apply gamma correction to convert from linear RGB to sRGB
	if r > 0.0031308 {
		r = 1.055*math.Pow(r, 1/2.4) - 0.055 // Apply gamma correction for values above 0.0031308
	} else {
		r = 12.92 * r // Apply linear correction for values below or equal to 0.0031308
	}

	if g > 0.0031308 {
		g = 1.055*math.Pow(g, 1/2.4) - 0.055 // Apply gamma correction for values above 0.0031308
	} else {
		g = 12.92 * g // Apply linear correction for values below or equal to 0.0031308
	}

	if linearB > 0.0031308 {
		linearB = 1.055*math.Pow(linearB, 1/2.4) - 0.055 // Apply gamma correction for values above 0.0031308
	} else {
		linearB = 12.92 * linearB // Apply linear correction for values below or equal to 0.0031308
	}

	// Clip values to the range [0, 1]
	r = math.Max(0, math.Min(1, r)) // Ensure red component is within valid range
	g = math.Max(0, math.Min(1, g)) // Ensure green component is within valid range
	b = math.Max(0, math.Min(1, linearB)) // Ensure blue component is within valid range

	return r, g, b
}