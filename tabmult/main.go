package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		n, err := strconv.Atoi(ar[1])
		if err != nil {
			z01.PrintRune('\n')
			return
		}

		for x := 1; x <= 9; x++ {
			mul := x * n
			res := itoa(x) + " x " + itoa(n) + " = " + itoa(mul) + "\n"
			for _, r := range []rune(res) {
				z01.PrintRune(r)
			}
		}

	} else {
		z01.PrintRune('\n')
	}

}

func itoa(n int) string {
	if n == 0 {
		return ""
	}
	last := n % 10
	return itoa(n/10) + string(rune(48+last))
}
