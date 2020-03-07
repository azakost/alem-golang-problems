package main

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func main() {
	// num1 := &NodeAddL{Num: 5}
	// num1 = pushFront(num1, 1)
	// num1 = pushFront(num1, 3)

	// num2 := &NodeAddL{Num: 2}
	// num2 = pushFront(num2, 9)
	// num2 = pushFront(num2, 5)

	// result := AddLinkedNumbers(num1, num2)
	// for tmp := result; tmp != nil; tmp = tmp.Next {
	// 	fmt.Print(tmp.Num)
	// 	if tmp.Next != nil {
	// 		fmt.Print(" -> ")
	// 	}
	// }
	// fmt.Println()
}

func pushFront(l *NodeAddL, num int) *NodeAddL {
	var tmp NodeAddL
	tmp.Num = num
	tmp.Next = l
	return &tmp
}

func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {
	n1 := listToInt(num1)
	n2 := listToInt(num2)
	sum := n1 + n2
	return intToList(sum)
}

func intToList(x int) *NodeAddL {
	var res *NodeAddL
	for x != 0 {
		res = pushFront(res, x%10)
		x /= 10
	}
	return res
}

func listToInt(list *NodeAddL) int {
	res := 0
	power := power(10, listLen(list))
	for it := list; it != nil; it = it.Next {
		power /= 10
		res = res + it.Num*power
	}
	return res
}

func power(a, b int) int {
	if b < 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	return a * power(a, b-1)
}

func listLen(l *NodeAddL) int {
	res := 0
	for it := l; it != nil; it = it.Next {
		res++
	}
	return res
}
