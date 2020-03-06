package main

import "fmt"

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}

func StrLen(str string) int {
	n := 0
	for range []rune(str) {
		n++
	}
	return n
}
