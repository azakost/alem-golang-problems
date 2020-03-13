package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {

}

func ReduceInt(f func(int, int) int, arr []int) {
	res := f(arr[0], arr[1])
	s := strconv.Itoa(res)
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
