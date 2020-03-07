package main

import (
	"fmt"
)

func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 0))
}

func Nauuo(plus, minus, rand int) string {
	if plus == minus {
		return "0"
	}

	f := func(a, b int, o string) string {
		if rand == 0 {
			return o
		} else {
			if a+rand >= b {
				return "?"
			}
			return o
		}
	}

	if plus > minus {
		return f(minus, plus, "+")
	} else {
		return f(plus, minus, "-")
	}
}
