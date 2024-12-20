package main

import (
	"fmt"
)

func main() {
	for _, n := range [...]int{1, 2, 3, 1, 2, 3, 1, 2, 3} {

		b := caller(n)
		if b {
			fmt.Printf("> than 2 :%d, %v\n", n, b)
		} else {
			fmt.Printf("< than 2 :%d, %v\n", n, b)
		}

	}
}

//less than 2 panic , returns default bool value
func caller(n int) bool {
	defer func() {
		if p := recover(); p != nil {
			//fmt.Printf("paniced.....with :%d\n", n)
		}
	}()
	callee(n)
	return true
}

func callee(n int) {
	if n < 3 {
		panic("n less than 2")
	}
}
