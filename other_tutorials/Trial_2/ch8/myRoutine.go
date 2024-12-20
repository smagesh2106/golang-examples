package main

import (
	"fmt"
	"time"
)

func main() {
	go Spinner(50 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("Fibbonacci (%d)  = %d\n", n, fibN)
}

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func fib2(n int) uint64 {
	var x, y uint64 = 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Printf("%v\t%v\n", x, y)
	}
	return x
}
