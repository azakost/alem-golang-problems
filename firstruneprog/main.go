package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Println(FirstRune(str))
}

func FirstRune(s string) rune {
	return []rune(s)[0]
}
