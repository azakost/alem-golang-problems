package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		for _, x := range reverse(ar[1]) {
			z01.PrintRune(x)
		}
	}
	z01.PrintRune('\n')
}

func reverse(s string) string {
	size := 1
	for _, r := range []rune(s) {
		if r == ' ' {
			size++
		}
	}
	arr := make([]string, size)
	word := ""
	n := size - 1
	for _, r := range []rune(s) {
		if r != ' ' {
			word += string(r)
		} else {
			arr[n] = word
			n--
			word = ""
		}
	}
	arr[n] = word
	res := ""
	for i, x := range arr {
		sp := " "
		if i == 0 {
			sp = ""
		}
		res += sp + x
	}

	return res
}
