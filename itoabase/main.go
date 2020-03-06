package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(11165156, 10))
}

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	ch := ""
	if value < 0 {
		value = -value
		ch = "-"
	}
	str := "0123456789ABCDEF"
	str = str[:base]
	l := len(str)
	txt := ""
	for x := 1; x < l; x++ {
		if value == 0 {
			break
		}
		txt = string(str[value%l]) + txt
		value = value / l
	}
	return ch + txt
}
