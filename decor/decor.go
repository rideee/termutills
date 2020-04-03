package decor

import (
	"fmt"
	"strings"
)

// Print - print input with ANSI ESC Control Sequence.
func Print(f string, b string, mode string, input string) {
	fmt.Print(MakeEscSeq(f, b, mode, input))
}

// Println - print input with ANSI ESC Control Sequence,
// and new line at the end. Same like fmt.Println
func Println(f string, b string, mode string, input string) {
	Print(f, b, mode, input)
	fmt.Println()
}

// Printf - print formatted input with ANSI ESC Control Sequence,
// same like fmt.Printf.
func Printf(f string, b string, mode string, format string, a ...interface{}) {
	fmt.Printf(MakeEscSeq(f, b, mode, format), a...)
}

// FgPrint - use decor.Print with only foreground param for formatting.
func FgPrint(f string, input string) {
	Print(f, "", "", input)
}

// FgPrintln - use decor.Println with only foreground param for formatting.
func FgPrintln(f string, input string) {
	Println(f, "", "", input)
}

// FgPrintf - use decor.Printf with only foreground param for formatting.
func FgPrintf(f string, format string, input ...interface{}) {
	Printf(f, "", "", format, input...)
}

// BgPrint - use decor.Print with only background param for formatting.
func BgPrint(b string, input string) {
	Print("", b, "", input)
}

// BgPrintln - use decor.Println with only background param for formatting.
func BgPrintln(b string, input string) {
	Println("", b, "", input)
}

// BgPrintf - use decor.Printf with only background param for formatting.
func BgPrintf(b string, format string, input ...interface{}) {
	Printf("", b, "", format, input...)
}

// MPrint - use decor.Print with only mode param for formatting.
func MPrint(mode string, input string) {
	Print("", "", mode, input)
}

// MPrintln - use decor.Println with only mode param for formatting.
func MPrintln(mode string, input string) {
	Print("", "", mode, input)
	fmt.Println()
}

// MPrintf - use decor.Printf with only mode param for formatting.
func MPrintf(mode string, format string, input ...interface{}) {
	Printf("", "", mode, format, input...)
}

// Return(string) input(string) with ANSI ESC Control Sequence.
func MakeEscSeq(f string, b string, mode string, input string) string {
	var combo string

	if f != "" || b != "" || mode != "" {
		combo = "\033["

		// Mode
		if mode != "" {
			combo += modeNr(mode)
		}

		// Foreground color
		if f != "" {
			if mode != "" {
				combo += ";"
			}

			fg, _ := coloNr(f)
			combo += fg

			if b != "" {
				combo += ";"
			}
		}

		// Background color
		if b != "" {
			_, bg := coloNr(b)
			combo += bg
		}

		combo += "m"
	}

	combo += input

	// Reset
	if f != "" || b != "" || mode != "" {
		fmt.Print("\033[0m")
		combo += "\033[0m"
	}

	return combo
}

// Helper func.
// coloNr return foreground and background color number (8/16).
func coloNr(v string) (string, string) {
	v = strings.ToLower(v)
	switch v {
	case "black":
		return "30", "40"
	case "red":
		return "31", "41"
	case "green":
		return "32", "42"
	case "yellow":
		return "33", "43"
	case "blue":
		return "34", "44"
	case "magenta":
		return "35", "45"
	case "cyan":
		return "36", "46"
	case "lgray":
		return "37", "47"
	case "dgray":
		return "90", "100"
	case "lred":
		return "91", "101"
	case "lgreen":
		return "92", "102"
	case "lyellow":
		return "93", "103"
	case "lblue":
		return "94", "104"
	case "lmagenta":
		return "95", "105"
	case "lcyan":
		return "96", "106"
	case "white":
		return "97", "107"
	default:
		return "39", "49"
	}
}

// Helper func.
// modeNr return formatting mode nr.
func modeNr(v string) string {
	v = strings.ToLower(v)
	switch v {
	case "b", "bold":
		return "1"
	case "d", "dim":
		return "2"
	case "u", "underline":
		return "4"
	case "bl", "blink":
		return "5"
	case "r", "reverse":
		return "7"
	default:
		return "0"
	}
}
