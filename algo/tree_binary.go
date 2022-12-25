package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (t *Tree) delete(data int) {
	if t == nil {
		return
	}

	if t.root.data == data {
		t.root = nil
	} else if data < t.root.data {

	} else if data > t.root.data {

	}
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &Tree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')

	fmt.Println("\n\n")

	tree.delete(100)
	print(os.Stdout, tree.root, 0, 'M')
}
