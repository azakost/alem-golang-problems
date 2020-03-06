package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		n, er := strconv.Atoi(ar[1])
		if er == nil {
			fprime(n)
		}
	}
	z01.PrintRune('\n')
}

func fprime(n int) {
	for x := 2; x < n*2; x++ {
		if n%x == 0 {
			fmt.Print(x)
			if n/x != 1 {
				fmt.Print("*")
			}
			fprime(n / x)
			break
		}
	}
}
