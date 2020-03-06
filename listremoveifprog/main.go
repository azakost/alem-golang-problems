package main

import "fmt"

func main() {
	var ls List
	ListPushBack(&ls, "a")
	ListPushBack(&ls, "b")
	ListPushBack(&ls, "c")

	fmt.Println("<<Before remove>>")
	node := ls.Head
	for node.Next != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println(node.Data)

	fmt.Println()
	fmt.Println("<<after remove>>")
	ListRemoveIf(&ls, "b")

	node2 := ls.Head
	for node2.Next != nil {
		fmt.Println(node2.Data)
		node2 = node2.Next
	}
	fmt.Println(node2.Data)

}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	var tmpList List
	for node := l.Head; node != nil; node = node.Next {
		if node.Data != data_ref {
			ListPushBack(&tmpList, node.Data)
		}
	}
	l.Head = tmpList.Head
}

func ListPushBack(l *List, d interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: d}
		return
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &NodeL{Data: d}
}
