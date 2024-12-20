package main

import (
	"fmt"
)

func main() {
	for _, n := range [...]int{1, 2, 3, 1, 2, 3, 1, 2, 3} {

		b := caller(n)
		if b < 3 {
			fmt.Printf(" <  than 3 :%d, %v\n", n, b)
		} else {
			fmt.Printf(" >= than 3 :%d, %v\n", n, b)
		}

	}
}

//less than 2 panic, returns default int val( 0 )
func caller(n int) int {
	defer func() {
		if p := recover(); p != nil {
			//fmt.Printf("paniced.....with :%d\n", n)
		}
	}()
	callee(n)
	return n
}

func callee(n int) {
	if n < 3 {
		panic("n less than 3")
	}
}
