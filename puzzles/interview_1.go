package main

import (
	"fmt"
)

func main() {
	find1()
}

func find1() {
	var found = make(map[int]bool)
	var index = make(map[int]int)

	var a = []int{2, 4, 5, 11, 12, 5}
	var find = 16

	for n, val := range a {

		req := find - val
		index[val] = n
		found[val] = true

		if _, ok := found[req]; ok {
			fmt.Printf("Indexs are %d, %d\n", n, index[req])
		}
	}
}
