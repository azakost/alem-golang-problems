package main

import (
	"fmt"
)

func main() {
	fmt.Println(Priorprime(14))
}

func Priorprime(x int) int {
	sum := 0
	for n := 2; n < x; n++ {
		if isPrime(n) {
			sum += n
		}
	}
	return sum
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
