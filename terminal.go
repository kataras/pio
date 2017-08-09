package pio

import (
	"io"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/kataras/pio/terminal"
)

func isTerminal(output io.Writer) bool {
	isTerminal := !IsNop(output) || terminal.IsTerminal(output)
	if isTerminal {
		return true
	}
	if runtime.GOOS != "windows" {
		return false
	}

	// check specific for windows operating system
	// versions, after windows 10 microsoft
	// gave suppprt for 256-color console.

	cmd := exec.Command("cmd", "ver")

	b, err := cmd.Output()
	if err != nil {
		return false
	}
	/*
		Microsoft Windows [Version 10.0.15063]
		(c) 2017 Microsoft Corporation. All rights reserved.
	*/
	lines := string(b)
	if lines == "" {
		return false
	}

	start := strings.IndexByte(lines, '[')
	end := strings.IndexByte(lines, ']')

	winLine := lines[start+1 : end]
	// Version 10.0.15063
	versionsLine := winLine[strings.IndexByte(winLine, ' ')+1:]
	// 10.0.15063
	versionSems := strings.Split(versionsLine, ".")
	// 10
	// 0
	// 15063
	if len(versionSems) < 3 {
		return false
	}

	// ok, we need to check if it's windows version 10
	if versionSems[0] != "10" {
		return false
	}

	buildNumber, err := strconv.Atoi(versionSems[2])
	// and the build number is equal or greater than 10586
	if err != nil {
		return false
	}

	return buildNumber >= 10586
}
