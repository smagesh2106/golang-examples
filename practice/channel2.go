package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	n := make(chan int)
	s := make(chan int)

	go counter(n)
	go square(s, n)
	printer(s)
}
