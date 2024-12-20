package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}
type root struct {
	node *node
}

func (r *root) Insert(data int) *root {
	if r.node == nil {
		r.node = &node{data: data, next: nil}
	} else {
		r.node.Insert(data)
	}
	return r
}

func (n *node) Insert(data int) {
	if n.next == nil {
		n.next = &node{data: data, next: nil}
	} else {
		n.next.Insert(data)
	}
}

func print(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("-> %v", n.data)
	print(n.next)
}
func main() {
	root := &root{}
	root.Insert(5).
		Insert(10).
		Insert(15)

	print(root.node)

}
