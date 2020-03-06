package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 3 {
		a, ea := strconv.Atoi(ar[1])
		b, eb := strconv.Atoi(ar[2])
		if ea == nil && eb == nil {
			fmt.Print(gcd(a, b))
		}
	}
	z01.PrintRune('\n')
}

func less(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	for x := less(a, b); x >= 0; x-- {
		if a%x == 0 && b%x == 0 {
			return x
		}
	}
	return 1
}
