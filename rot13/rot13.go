package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	// Rot13
	ar := os.Args
	if len(ar) == 2 {
		for _, r := range ar[1] {
			if r >= 'a' && r <= 'z' {
				if (r + 13) > 'z' {
					k := (r + 12) - 'z'
					z01.PrintRune('a' + k)
				} else {
					z01.PrintRune(r + 13)
				}
			} else if r >= 'A' && r <= 'Z' {
				if (r + 13) > 'Z' {
					k := (r + 12) - 'Z'
					z01.PrintRune('A' + k)
				} else {
					z01.PrintRune(r + 13)
				}
			} else {
				z01.PrintRune(r)
			}
		}
	}
	z01.PrintRune('\n')
}
