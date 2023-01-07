package main

import (
	"fmt"
)

func recur2(i int) {
	if i == 0 {
		//fmt.Println(i)
		fmt.Println("+++>", i)
		return
	}
	for j := i; j > 0; j-- {
		fmt.Println("--->", i)
		recur2(i - 1)
		fmt.Println("===>", i)
		//recur(i - 1)
	}
}

func main() {
	recur2(2)
}
