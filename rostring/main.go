package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	ar := os.Args

	if len(ar) == 2 {
		for _, r := range []rune(rostring(ar[1])) {
			z01.PrintRune(r)
		}

	}
	z01.PrintRune('\n')
}

func clear(s string) string {
	if len(s) == 0 {
		return s
	}

	if s[0] == ' ' {
		return clear(s[1:])
	}

	l := len(s) - 1
	if s[l] == ' ' {
		return clear(s[:l])
	}

	return s
}

func midsp(s string) (string, int) {
	res := ""
	flag := true
	n := 1
	for _, x := range []rune(clear(s)) {
		if x != ' ' {
			res += string(x)
			flag = true
		} else {
			if flag {
				res += string(x)
				n++
				flag = false
			}
		}
	}
	return res, n
}

func rostring(s string) string {
	str, size := midsp(s)
	if size == 1 {
		return str
	}
	arr := make([]string, size)
	word := ""
	n := 0
	k := 0
	tmp := ""
	for _, x := range []rune(str) {
		if x != ' ' {
			word += string(x)
		} else {
			if k == 0 {
				tmp = word
				word = ""
				k++
			} else {
				arr[n] = word
				n++
				word = ""
			}
		}
	}
	arr[n] = word
	arr[n+1] = tmp

	new := ""
	for i, x := range arr {
		s := " "
		if i == 0 {
			s = ""
		}
		new += s + x
	}

	return new
}
