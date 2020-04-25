package pio

import (
	"fmt"
)

func ExampleRich() {
	fmt.Println(Rich("black", Black))
	fmt.Println(Rich("cyan", Cyan))

	fmt.Println(Rich("background cyan",
		Cyan,
		Background,
	))

	fmt.Println(Rich("bright yellow",
		Yellow,
		Bright,
	))

	fmt.Println(Rich("bright bold magenta",
		Magenta,
		Bright,
		Bold,
	))

	fmt.Println(Rich("bold cyan",
		Cyan,
		Bold,
	))

	fmt.Println(Rich("bold underline reversed bright cyan",
		Cyan,
		Bright,
		Bold,
		Underline,
		Reversed,
	))

	fmt.Println(Rich("extended 256-color custom color: 153 (blue-ish)", 153))
	fmt.Println(Rich("extended 256-color custom bright reversed color: 153 (blue-ish)", 153,
		Bright,
		Reversed,
	))

	// Output:
	// [30mblack[0m
	// [36mcyan[0m
	// [46mbackground cyan[0m
	// [33;1mbright yellow[0m
	// [35;1m[1mbright bold magenta[0m
	// [36m[1mbold cyan[0m
	// [36;1m[7m[4m[1mbold underline reversed bright cyan[0m
	// [38;5;153mextended 256-color custom color: 153 (blue-ish)[0m
	// [38;5;153;1m[7mextended 256-color custom bright reversed color: 153 (blue-ish)[0m
}
