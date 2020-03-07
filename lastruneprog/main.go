package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Println(LastRune(str))
}

func LastRune(s string) rune {
	var res rune
	for _, v := range s {
		res = v
	}
	return res
}
