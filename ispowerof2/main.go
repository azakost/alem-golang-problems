package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		n, ea := strconv.Atoi(ar[1])

		if ea != nil {
			z01.PrintRune('\n')
			return
		}

		if isPowerofTwo(n) {
			for _, r := range "true" {
				z01.PrintRune(r)
			}
		} else {
			for _, r := range "false" {
				z01.PrintRune(r)
			}
		}
	}
	z01.PrintRune('\n')
}

// func isPowerofTwo(n int) bool {
// 	if n == 0 {
// 		return false
// 	}
// 	z := n % 2
// 	if z == 0 {
// 		return isPowerOf2(n / 2)
// 	}
// 	if n == 1 {
// 		return true
// 	}
// 	return false
// }

func isPowerofTwo(n int) bool {
	if n&(n-1) == 0 { // Конъюнкция
		return true
	}
	return false
}
