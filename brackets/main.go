package main

import (
	"fmt"
	"os"
)

func main() {
	ar := os.Args
	if len(ar) > 1 {
		for _, x := range ar[1:] {
			fmt.Println(brackets(x))
		}
	} else {
		fmt.Println()
	}
}

func brackets(s string) string {
	if !isOK('(', ')', s) {
		return "Error"
	}
	if !isOK('[', ']', s) {
		return "Error"
	}
	if !isOK('{', '}', s) {
		return "Error"
	}
	return "OK"
}

func isOK(a, b rune, s string) bool {
	if s == "" {
		return true
	}
	ac := 0
	bc := 0
	for _, x := range []rune(s) {
		if x == a {
			ac++
		}
		if x == b {
			bc++
		}
	}
	if ac == bc {
		return true
	}
	return false
}
