package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 3 {
		a, ea := strconv.Atoi(ar[1])
		if a <= 0 || ea != nil {
			printer("Error\n")
			return
		}
		b, eb := strconv.Atoi(ar[2])
		if b <= 0 || eb != nil {
			printer("Error\n")
			return
		}
		board := chess(a, b)
		printer(board)
	} else {
		printer("Error\n")
	}
}

func chess(a, b int) string {
	res := ""
	for x := 1; x <= b; x++ {
		n := 0
		for y := 1; y <= a; y++ {
			if (x+y)%2 == 0 {
				res += "#"
			} else {
				res += " "
			}
			n++
			if n == a {
				res += "\n"
			}
		}
	}
	return res

}

func printer(s string) {
	for _, r := range []rune(s) {
		z01.PrintRune(r)
	}
}
