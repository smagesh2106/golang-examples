// threads.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	go Spinner(time.Millisecond * 200)
	fibN := fib(50)
	fmt.Printf("\rFibnonacci :%d\n", fibN)
	fmt.Println("Done")
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
