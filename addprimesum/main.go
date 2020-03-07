package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ar := os.Args
	if len(ar) == 2 {
		for _, x := range addPrime(ar[1]) {
			z01.PrintRune(x)
		}
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}

func addPrime(num string) string {
	n := atoi(num)
	sum := 0
	for x := 2; x <= n; x++ {
		if isPrime(x) {
			sum += x
		}
	}
	return itoa(sum)
}

func isPrime(n int) bool {
	z := 0
	for x := 2; x <= n; x++ {
		if n%x == 0 {
			z++
		}
	}
	if z == 1 {
		return true
	}
	return false
}

func atoi(s string) int {
	n := 0
	if s[0] == '-' {
		return 0
	}
	for _, x := range []byte(s) {
		x -= '0'
		n = n*10 + int(x)
	}
	return n
}

func itoa(n int) string {
	digits := "0123456789"
	if n < 10 {
		return digits[n : n+1]
	}
	return itoa(n/10) + digits[n%10:n%10+1]
}
