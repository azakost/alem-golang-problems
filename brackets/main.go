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
	br := ""
	for _, r := range s {
		if (r == ')' || r == '}' || r == ']') && len(br) == 0 {
			return "Error"
		}

		if r == '(' || r == '{' || r == '[' {
			br += string(r)
		}

		if r == ')' && br[len(br)-1] == '(' {
			br = br[:len(br)-1]
		}
		if r == '}' && br[len(br)-1] == '{' {
			br = br[:len(br)-1]
		}
		if r == ']' && br[len(br)-1] == '[' {
			br = br[:len(br)-1]
		}
	}
	if len(br) == 0 {
		return "OK"
	}
	return "Error"
}
