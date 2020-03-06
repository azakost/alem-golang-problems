package main

func pushFront(node *NodeAddL, num int) *NodeAddL {
	var tmp NodeAddL
	tmp.Num = num
	tmp.Next = node
	return &tmp
}

func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num: 5}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 3)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num: 2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 5)

	// 9 -> 0 -> 7
	// result := AddLinkedNumbers(num1, num2)
	// for tmp := result; tmp != nil; tmp = tmp.Next {
	// 	fmt.Print(tmp.Num)
	// 	if tmp.Next != nil {
	// 		fmt.Print(" -> ")
	// 	}
	// }
	// fmt.Println()

}

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

// func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {

// }
