package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		printer(robot(ar[1]))
	}

	z01.PrintRune('\n')
}

func robot(s string) string {
	x := 0
	y := 0
	for _, r := range []rune(s) {
		if r == 'U' {
			y++
		}
		if r == 'D' {
			y--
		}
		if r == 'R' {
			x++
		}
		if r == 'L' {
			x--
		}
	}

	if x == 0 && y == 0 {
		return "true"
	}
	return "false"
}

func printer(s string) {
	for _, r := range []rune(s) {
		z01.PrintRune(r)
	}
}
