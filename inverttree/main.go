package main

func main() {

}

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

func InvertTree(root *TNode) *TNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
