package main

import "fmt"

func main() {
	fmt.Println(fib1(6))
}

func fib1(n int) int {

	if n == 1 || n == 2 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}
