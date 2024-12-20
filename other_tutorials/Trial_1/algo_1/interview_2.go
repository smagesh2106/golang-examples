package main

import (
	"fmt"
)

func main() {
	type dat struct {
		n   int
		val bool
	}
	var b = make(map[int]dat)
	//	var a = []int{17, 2, 4, 5, 11, 5, 1, 7, 18}
	//	var find = 18
	var a = []int{2, 4, 5, 11, 12, 5}
	var find = 16

	for n, val := range a {
		b[val] = dat{n: n, val: true}

		req := find - val
		_, ok := b[req]
		if ok && b[req].val {
			fmt.Printf("Indexs are %d, %d\n", n, b[req].n)
		}
	}
}
