package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		for _, x := range ar[1] {
			for z := 0; z < times(x); z++ {
				z01.PrintRune(x)
			}
		}
	}

	z01.PrintRune('\n')
}

func times(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	} else if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 1
	} else {
		return 1
	}
}
