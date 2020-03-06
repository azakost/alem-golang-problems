package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 3 {
		for _, r := range []rune(clear(join(ar[1:]))) {
			z01.PrintRune(r)
		}
	}

	z01.PrintRune('\n')
}

func join(arr []string) string {
	res := ""
	for _, x := range arr {
		res += x
	}
	return res
}

func has(r rune, s string) bool {
	for _, x := range []rune(s) {
		if x == r {
			return true
		}
	}
	return false
}

func clear(s string) string {
	z := ""
	for _, x := range []rune(s) {
		if !has(x, z) {
			z += string(x)
		}
	}
	return z
}
