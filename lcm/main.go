package main

import "fmt"

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}

func Lcm(first, second int) int {
	for x := 2; x < first*second; x++ {
		if first%x == 0 && second%x == 0 {
			return x
		}
	}
	return first * second
}
