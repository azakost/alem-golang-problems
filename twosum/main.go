package main

import "fmt"

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 7)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {
	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			x := nums[a]
			y := nums[b]
			if x+y == target {
				return []int{a, b}
			}
		}
	}
	return []int{}
}
