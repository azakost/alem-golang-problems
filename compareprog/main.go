package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 3 {
		for _, r := range Compare(ar[1], ar[2]) {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func Compare(a, b string) string {
	if a == b {
		return "0"
	} else if a < b {
		return "-1"
	} else {
		return "1"
	}
}
