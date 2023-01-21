package main

import (
	"fmt"
)

func recur7(i int) {
	if i == 0 {
		//fmt.Println(i)
		fmt.Println("+++>", i)
		return
	}
	for j := 0; j < i; j++ {
		fmt.Printf("--->i = %d, j = %d\n", i, j)
		recur7(i - 1)
		fmt.Printf("===>i = %d, j = %d\n", i, j)
	}
	//recur(i - 1)
}

func main() {
	recur7(3)
}
