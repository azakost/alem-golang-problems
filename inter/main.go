package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 3 {
		for _, x := range []rune(inter(ar[1:])) {
			z01.PrintRune(x)
		}
	}
	z01.PrintRune('\n')
}

func inter(arr []string) string {
	a := clear(arr[0])
	b := arr[1]
	res := ""
	for _, x := range []rune(a) {
		if has(x, b) {
			res += string(x)
		}
	}
	return res
}

func clear(s string) string {
	res := ""
	for _, r := range []rune(s) {
		if !has(r, res) {
			res += string(r)
		}
	}
	return res
}

func has(r rune, s string) bool {
	for _, x := range []rune(s) {
		if r == x {
			return true
		}
	}
	return false
}
