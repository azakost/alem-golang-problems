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
	num, opr := twoSlices(s)
	if len(num)-1 != len(opr) {
		return "Error"
	}

	res := 0
	for x := 0; x < len(num); x++ {
		a, e := strconv.Atoi(num[x])
		if e != nil {
			return "Error"
		}
		if x == 0 {
			res = a
		} else {
			res = calc(res, opr[x-1], a)
		}
	}
	return itoa(res)
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

func twoSlices(s string) ([]string, []string) {
	s = clean(s)
	arr := strings.Split(s, " ")
	var opr []string
	var dig []string
	for _, r := range arr {
		if r == "+" || r == "-" || r == "*" || r == "/" || r == "%" {
			opr = append(opr, r)
		} else {
			dig = append(dig, r)
		}
	}
	return dig, opr
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
	for _, r := range []rune(s) {
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
