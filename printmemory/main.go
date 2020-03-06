package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func PrintMemory(arr [10]int) {
	str := ""
	txt := ""
	z := 0
	for _, n := range arr {
		h := hex(n)
		if len(h) < 4 {
			for x := 0; x <= 5-len(h); x++ {
				h += "0"
			}
		}
		h += " 0000"
		if len(str) == 0 || z == 0 {
			str += h
		} else {
			str += " " + h
		}
		z++
		if z == 4 {
			str += "\n"
			z = 0
		}
		if n < 32 {
			txt += string(rune('.'))
		} else {
			txt += string(rune(n))
		}
	}
	println(str)
	println(txt)
}

func hex(n int) string {
	base := "0123456789abcdef"
	if n < 16 {
		return base[n : n+1]
	}
	return hex(n/16) + string(base[n%16:n%16+1])
}

func println(s string) {
	for _, r := range []rune(s) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
