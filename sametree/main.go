package main

type TreeNodeL struct {
	Left  *TreeNodeL
	Val   int
	Right *TreeNodeL
}

func main() {

}

func IsSameTree(p *TreeNodeL, q *TreeNodeL) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	l := IsSameTree(p.Left, q.Left)
	r := IsSameTree(p.Right, q.Right)
	if l && r {
		return true
	}
	return false
}
