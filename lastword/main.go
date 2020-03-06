package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		for _, r := range []rune(div(ar[1])) {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func div(s string) string {
	l := len(s) - 1
	res := ""
	if s[l] == ' ' {
		return div(s[:l])
	} else {
		for x := l; x >= 0; x-- {
			if s[x] != ' ' {
				res = string(s[x]) + res
			} else {
				return res
			}
		}
	}
	return res
}
