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

func fun(n int) {
	if n > 2 {
		fun(n - 1)
		fun(n - 2)
		fun(n - 3)
	}
	fmt.Println(n)
}
func main() {
	recur(10)
	fun(5)
}
