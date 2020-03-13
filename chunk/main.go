package main

import "fmt"

func main() {

}

func Chunk(slice []int, size int) {
	l := Len(slice)
	if size <= 0 {
		fmt.Println()
		return
	}
	if l == 0 {
		fmt.Println([][]int{})
		return
	}
	msize := l / size
	if l%size != 0 {
		msize++
	}
	mother := make([][]int, msize)
	child := make([]int, size)
	c := 0
	m := 0
	for i, x := range slice {
		if c < size {
			child[c] = x
		} else {
			mother[m] = child
			if l-i < size {
				size = l - i
			}
			child = make([]int, size)
			c = 0
			child[c] = x
			m++
		}
		c++
	}
	mother[m] = child
	fmt.Println(mother)
}

func Len(ar []int) int {
	n := 0
	for range ar {
		n++
	}
	return n
}
