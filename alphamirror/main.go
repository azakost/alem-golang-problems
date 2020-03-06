package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		s := ar[1]
		for _, r := range []rune(s) {
			var res rune
			switch caze(r) {
			case "up":
				res = result(r, 'A', 'N', 'Z')
			case "lo":
				res = result(r, 'a', 'n', 'z')
			case "ch":
				z01.PrintRune(r)
			}
			z01.PrintRune(res)
		}
	}
	z01.PrintRune('\n')
}

func result(r, a, m, e rune) rune {
	if r < m {
		return e - r + a
	} else {
		return a + e - r
	}
}

func caze(r rune) string {
	if r >= 'A' && r <= 'Z' {
		return "up"
	} else if r >= 'a' && r <= 'z' {
		return "lo"
	}
	return "ch"
}
