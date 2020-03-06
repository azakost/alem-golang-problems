package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if !valid(base) || len(base) < 2 {
		return 0
	}
	pow := 1
	num := 0
	for x := len(s) - 1; x >= 0; x-- {
		idx := index(rune(s[x]), base)
		num += idx * pow
		pow *= len(base)
	}
	return num
}

func index(r rune, base string) int {
	for i, x := range []rune(base) {
		if x == r {
			return i
		}
	}
	return 0
}

func valid(base string) bool {
	tmp := ""
	for _, r := range []rune(base) {
		if r == '+' || r == '-' || inString(r, tmp) {
			return false
		}
		tmp += string(r)
	}
	return true
}

func inString(r rune, s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, x := range []rune(s) {
		if x == r {
			return true
		}
	}
	return false
}
