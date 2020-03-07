package main

import (
	"fmt"
	"os"
)

var chars = "abcdefghijklmnopqrstuvwxyz"

func main() {
	ar := os.Args
	if len(ar) >= 1 {
		fmt.Println(options(ar[1:]))
	}
}

func options(ar []string) string {
	if len(ar) == 0 {
		return "options: " + chars
	}

	new := ""
	for _, z := range ar {
		if z[:2] == "-h" {
			return "options: " + chars
		}
		for i, r := range z {
			if i == 0 {
				continue
			}
			if r < 'a' || r > 'z' {
				return "Invalid Option"
			}
		}
		new += z
	}
	k := 1
	str := "000000"
	for _, r := range rev(chars) {
		n := 0
		for _, k := range new {
			if r == k {
				n++
				break
			}
		}
		if len(str)-(k-1) == 8*k {
			str += " "
			k++
		}
		if n > 0 {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func rev(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	return s[l-1:] + rev(s[:l-1])
}
