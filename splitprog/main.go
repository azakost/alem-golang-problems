package main

import "fmt"

func main() {
	str := "HelloHAHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func Split(str, charset string) []string {
	st, size := div(str, charset, 1)
	word := ""
	res := make([]string, size)
	n := 0
	for _, x := range st {
		if x != '\v' {
			word += string(x)
		} else {
			res[n] = word
			n++
			word = ""
		}
	}
	res[n] = word
	return res
}

func div(s, c string, n int) (string, int) {
	cl := len(c)
	sl := len(s)
	if sl < cl {
		return s, n
	}

	if s[0:cl] == c {
		return div("\v"+s[cl:], c, n+1)
	}
	st, k := div(s[1:], c, n)
	return s[0:1] + st, k
}
