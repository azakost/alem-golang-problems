package main

import (
	"fmt"
	"os"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		fmt.Print(balanced(ar[1]))
	}
	fmt.Println()
}

func balanced(s string) int {
	n := 0
	k := 0
	for i, r := range s {
		if r == 'C' {
			n++
		}
		if r == 'D' {
			n--
		}
		if n == 0 && i != 0 {
			k++
		}
	}
	return k
}
