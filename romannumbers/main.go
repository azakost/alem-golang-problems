package main

import (
	"fmt"
	"os"
)

func main() {
	ar := os.Args
	if Len(ar) == 2 {
		n, valid := atoi(ar[1])
		if !valid {
			fmt.Printf("ERROR: can not convert to roman digit\n")
			return
		}
		fmt.Printf(toRoman(n))
	}
}

func toRoman(n int) string {
	res := ""
	if n < 0 {
		n = -n
		res = "-"
	}
	num := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	sym := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	m := make(map[int]string)
	i := 12
	z := 0
	for n > 0 {
		div := n / num[i]
		n %= num[i]
		for div > 0 {
			m[z] = sym[i]
			res += sym[i]
			div--
			z++
		}
		i--
	}
	calc := ""
	for x := 0; x < z; x++ {
		if Len(m[x]) == 1 {
			calc += m[x]
		} else {
			calc += "(" + string(m[x][1]) + "-" + string(m[x][0]) + ")"
		}
		if x != z-1 {
			calc += "+"
		}
	}
	if res[0] == '-' {
		calc = "-(" + calc + ")"
	}
	return calc + "\n" + res + "\n"
}

func Len(s interface{}) int {
	n := 0
	switch s.(type) {
	case string:
		for range s.(string) {
			n++
		}
	case []string:
		for range s.([]string) {
			n++
		}
	}
	return n
}

func atoi(s string) (int, bool) {
	if !valid(s) {
		return 0, false
	}
	m := 1
	if s[0] == '+' {
		s = s[1:]
	}
	if s[0] == '-' {
		s = s[1:]
		m *= -1
	}
	n := 0
	for _, x := range []byte(s) {
		x -= '0'
		n = n*10 + int(x)
	}
	if m*n > 4000 || m*n == 0 {
		return m * n, false
	}
	return m * n, true
}

func valid(s string) bool {
	for i, r := range s {
		if i == 0 && (r == '-' || r == '+') {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
