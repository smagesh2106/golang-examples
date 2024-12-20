package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
}

func (L *List) insertNode(data int) {
	n := Node{value: data, next: nil}
	ptr := L.head
	if ptr == nil {
		L.head = &n
	} else {
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = &n
	}
}

func main() {
	var L = &List{}
	for i := 0; i < 10; i++ {
		L.insertNode(i)
	}

	//ptr := L.head
	/*
		for ptr.next != nil {
			fmt.Printf("%v->%d\n", ptr, ptr.value)
			ptr = ptr.next
		}
		fmt.Printf("%v->%d\n", ptr, ptr.value)
	*/
	display(L.head)
	L.insertAfter(0, 20)
	display(L.head)
	L.insertBefore(3, 30)
	display(L.head)
	L.insertAfter(9, 40)
	display(L.head)

	fmt.Println("\n\n")
	//L.delete(0)
	//display(L.head)
	//L.insertBefore(0, 21)
	//display(L.head)
}

func display(n *Node) {
	fmt.Printf("->:%v ", n.value)
	if n.next == nil {
		fmt.Println("\n")
		return
	} else {
		display(n.next)
	}
}

/*
func (L *List) insertAfter(marker, data int) {
	n := Node{value: data, next: nil}
	ptr := L.head

	for ptr.next != nil {
		if ptr.value == marker {
			break
		}
		ptr = ptr.next
	}

	if ptr.next != nil {
		tmp := ptr.next
		n.next = tmp
		ptr.next = &n
	} else {
		ptr.next = &n
	}
}
*/

func (L *List) insertAfter(marker, data int) {
	n := Node{value: data, next: nil}
	ptr := L.head

	if L.head.value == marker {
		n.next = L.head.next
		L.head.next = &n
		return
	}

	found := false
	for ; ptr.next != nil; ptr = ptr.next {
		if ptr.value == marker {
			n.next = ptr.next
			ptr.next = &n
			found = true
			break
		}
	}

	if !found && ptr.next == nil {
		n.next = nil
		ptr.next = &n
	}
}

func (L *List) insertBefore(marker, data int) {
	n := Node{value: data, next: nil}
	ptr := L.head

	if L.head.value == marker {
		n.next = L.head
		L.head = &n
		return
	}

	for p := ptr; p.next != nil; p = p.next {
		if p.next.value == marker {
			n.next = p.next
			p.next = &n
			break
		}
	}
}

func (L *List) delete(data int) {

	ptr := L.head

	//first node
	if L.head.value == data {
		L.head = L.head.next
		return
	}

	for p := ptr; p.next != nil; p = p.next {
		//last node
		if p.next.value == data {
			if p.next.next == nil {
				p.next = nil
				break
			} else {
				tmp := p.next.next
				p.next = tmp
				break
			}
		}
	}
}
