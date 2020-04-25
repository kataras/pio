package main

import (
	"os"

	"github.com/kataras/pio"
)

func main() {
	// pio will automatically detect the operating system/terminal which
	// is running under and will provide color support.
	p := pio.NewTextPrinter("color", os.Stdout)
	p.Println(pio.Rich("this is a blue text", pio.Blue))
	p.Println(pio.Rich("this is a gray text", pio.Gray))
	p.Println(pio.Rich("this is a red text", pio.Red))
	p.Println(pio.Rich("this is a purple text", pio.Magenta))
	p.Println(pio.Rich("this is a yellow text", pio.Yellow))
	p.Println(pio.Rich("this is a green text", pio.Green))
}
