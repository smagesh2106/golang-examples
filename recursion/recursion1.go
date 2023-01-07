package main

import (
	"fmt"
)

func recur(i int) {
	if i == 0 {
		//fmt.Println(i)
		fmt.Println("+++>", i)
		return
	}
	fmt.Println("--->", i)
	recur(i - 1)
	fmt.Println("===>", i)
	//recur(i - 1)
}

func main() {
	recur(3)
}
