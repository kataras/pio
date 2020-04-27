package pio

import (
	"io"
	"runtime"

	"github.com/kataras/pio/terminal"
)

// SupportColors reports whether the "w" io.Writer is not a file and it does support colors.
func SupportColors(w io.Writer) bool {
	isTerminal := !IsNop(w) && terminal.IsTerminal(w)
	if isTerminal && runtime.GOOS == "windows" {
		// if on windows then return true only when it does support 256-bit colors,
		// this is why we initially do that terminal check for the "w" writer.
		return terminal.SupportColors
	}

	return isTerminal
}
