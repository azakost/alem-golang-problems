package main

import "fmt"

func main() {
	var list List
	pushBack(&list, &NodeL{Data: 1})
	pushBack(&list, &NodeL{Data: 1})
	fmt.Print(ListSize(&list))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func pushBack(l *List, d *NodeL) {
	if l.Head == nil {
		l.Head = d
		return
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = d
}

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}

	node := l.Head
	n := 1
	for node.Next != nil {
		node = node.Next
		n++
	}
	return n
}
