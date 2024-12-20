package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func main() {
	n1 := new(Node)
	n1.value = 7
	n1.next = nil

	n2 := new(Node)
	n2.value = 8
	n2.next = n1

	n3 := new(Node)
	n3.value = 9
	n3.next = n2

	f1(n3)
}

func f1(n *Node) {
	if n.next == nil {
		fmt.Printf("%v", n.value)
		return
	} else {
		fmt.Printf("%v->", n.value)
		f1(n.next)
	}
}
