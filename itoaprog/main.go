package main

import "fmt"

func main() {
	fmt.Println(Itoa(100 - 100))
}

func Itoa(n int) string {
	num := n
	ch := ""
	if n < 0 {
		num = -n
		ch = "-"
	}
	digits := "0123456789"
	if num < 10 {
		return ch + digits[num:num+1]
	}
	return ch + Itoa(num/10) + digits[num%10:num%10+1]
}
