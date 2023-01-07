package main

import "fmt"

var cache2 = [9]int64{}

func fib2(n int64) int64 {
	if n == 2 || n == 1 {
		return 1
	}
	if cache2[n] != 0 {
		return cache2[n]
	}
	cache2[n] = fib2(n-1) + fib2(n-2)
	fmt.Printf(" %d", n)
	return cache2[n]

}

func main() {
	fmt.Println(fib2(8))
}
