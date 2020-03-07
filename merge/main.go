package main

import "fmt"

func main() {

}

type TreeNodeM struct {
	Left  *TreeNodeM
	Val   int
	Right *TreeNodeM
}

func MergeTrees(t1 *TreeNodeM, t2 *TreeNodeM) *TreeNodeM {

	newRoot := &TreeNodeM{}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	newRoot.Val = t1.Val + t2.Val
	newRoot.Left = MergeTrees(t1.Left, t2.Left)
	newRoot.Right = MergeTrees(t1.Right, t2.Right)
	return newRoot
}

func PrintAnotherTree(root *TreeNodeM) {
	if root != nil {
		PrintAnotherTree(root.Left)
		fmt.Println(root.Val)
		PrintAnotherTree(root.Right)
	}
}

func InsertAnotherTNode(root *TreeNodeM, data int) *TreeNodeM {
	newNode := &TreeNodeM{Val: data}
	if root == nil {
		return newNode
	}
	if data < root.Val {
		root.Left = InsertAnotherTNode(root.Left, data)
	} else {
		root.Right = InsertAnotherTNode(root.Right, data)
	}
	return root
}
