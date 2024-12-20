package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 10
	//fibN := fib(n)
	//fmt.Printf("\r Fibonacci-1(%d) : %d\n", n, fibN)
	fibN2 := fib2(n)
	fmt.Printf("\n\r Fibonacci-2(%d) : %d\n", n, fibN2)

}

func spinner(delay time.Duration) {
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

func fib2(n int) int {
	x, y := 0, 1

	fmt.Printf("%d ", x)
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Printf("%d ", x)
	}
	return x
}
