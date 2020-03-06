package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		printer(compare(ar[1]))
	}
	z01.PrintRune('\n')

}

func compare(s string) string {
	arr := count(s)
	for _, x := range arr {
		if inArray(x, arr) > 1 {
			return "false"
		}
	}
	return "true"
}

func inArray(n int, arr []int) int {
	z := 0
	for _, x := range arr {
		if x == n {
			z++
		}
	}
	return z
}

func count(s string) []int {
	c := formUniq(s)
	nums := make([]int, len(c))
	m := 0
	for _, a := range []rune(c) {
		n := 0
		for _, b := range []rune(s) {
			if a == b {
				n++
			}
		}
		nums[m] = n
		m++
	}
	return nums
}

func formUniq(s string) string {
	new := ""
	for _, a := range []rune(s) {
		if isUniq(a, new) {
			new += string(a)
		}
	}
	return new
}

func isUniq(r rune, s string) bool {
	for _, a := range []rune(s) {
		if r == a {
			return false
		}
	}
	return true
}

func printer(s string) {
	for _, r := range []rune(s) {
		z01.PrintRune(r)
	}
}
