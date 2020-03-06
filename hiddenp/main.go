package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 3 {
		fmt.Print(hiddenp(ar[1], ar[2]))
	}
	z01.PrintRune('\n')

}

func hiddenp(a, b string) int {

	if len(a) == 0 {
		return 1
	}

	match := 0
	x := 0
	y := 0
	for x < len(a) {
		for y < len(b) {
			if a[x] == b[y] {
				match++
				x++
			}
			if match == len(a) {
				return 1
			}
			y++
		}
		x++
	}
	return 0
}
