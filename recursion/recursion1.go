package main

import (
	"fmt"
)

func recur(i int) {
	if i == 0 {
		//fmt.Println(i)
		fmt.Println("==================")
		return
	}
	fmt.Println("--->", i)
	recur(i - 1)
	fmt.Println("===>", i)
}

func main() {
	recur(10)
}
