package main

import (
	"fmt"
)

func fun(n int) {
	if n > 2 {
		fun(n - 1)
		fun(n - 2)
		fun(n - 3)
	}
	fmt.Println(n)
}
func main() {
	fun(5)
}
