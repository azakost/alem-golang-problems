package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ar := os.Args

	if len(ar) == 4 {
		a, ea := strconv.Atoi(ar[1])
		b, eb := strconv.Atoi(ar[3])
		op := ar[2]
		if ea != nil || eb != nil {
			fmt.Println(0)
		} else {
			if op == "/" && b == 0 {
				fmt.Println("No division by 0")
			} else if op == "%" && b == 0 {
				fmt.Println("No modulo by 0")
			} else {
				fmt.Println(calc(a, b, op))
			}

		}
	}
}

func calc(a, b int, op string) int {
	var sum int
	var check int

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	switch op {
	case "+":
		if a > 0 && b > maxInt-a {
			return 0
		}
		if a <= 0 && b < minInt-a {
			return 0
		}
		return a + b

	case "-":
		if a < 0 && b > 0 {
			if a < minInt+b {
				return 0
			}
		}
		if a > 0 && b < 0 {
			if a > maxInt+b {
				return 0
			}
		}
		return a - b

	case "*":
		sum = a * b
		check = sum / a

	case "/":
		sum = a / b
		check = sum * a

	case "%":
		return a % b
	}

	if check == b {
		return sum
	}
	return 0
}
