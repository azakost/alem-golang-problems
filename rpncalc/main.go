package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		fmt.Println(rpncalc(ar[1]))
	} else {
		fmt.Println("Error")
	}
}

func rpncalc(s string) string {
	s = clean(s)
	arr := strings.Split(s, " ")
	var stack []int
	for i, x := range arr {

		if i < 2 && isOpr(x) {
			return "Error"
		}

		if !isOpr(x) {
			n, e := strconv.Atoi(x)
			if e != nil {
				return "Error"
			}
			stack = append(stack, n)
			continue
		}

		if isOpr(x) {
			l := len(stack)
			a := stack[l-2]
			b := stack[l-1]
			res := calc(a, x, b)
			stack = stack[:l-2]
			stack = append(stack, res)
			continue
		}
	}
	if len(stack) != 1 {
		return "Error"
	}
	return itoa(stack[0])
}

func isOpr(r string) bool {
	opr := "+-*/%"
	for _, x := range opr {
		if string(x) == r {
			return true
		}
	}
	return false
}

func calc(a int, o string, b int) int {
	s := 0
	switch o {
	case "+":
		s = a + b
	case "-":
		s = a - b
	case "*":
		s = a * b
	case "/":
		s = a / b
	case "%":
		s = a % b
	}
	return s
}

func clean(s string) string {
	if s[0] == ' ' {
		return clean(s[1:])
	}
	l := len(s) - 1
	if s[l] == ' ' {
		return clean(s[:l])
	}
	txt := ""
	flag := true
	for _, r := range s {
		if r != ' ' {
			txt += string(r)
			flag = true
		}
		if r == ' ' && flag {
			txt += " "
			flag = false
		}
	}
	return txt
}

func itoa(n int) string {
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
	return ch + itoa(num/10) + digits[num%10:num%10+1]
}
