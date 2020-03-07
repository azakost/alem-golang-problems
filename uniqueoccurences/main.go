package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		printer(uniq(ar[1]))
	}
	z01.PrintRune('\n')
}

func uniq(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		if _, ok := m[r]; ok {
			m[r]++
		} else {
			m[r] = 1
		}
	}
	n := make(map[int]int)
	for _, i := range m {
		if _, ok := n[i]; ok {
			return "false"
		} else {
			n[i] = 1
		}
	}
	return "true"
}

func printer(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
