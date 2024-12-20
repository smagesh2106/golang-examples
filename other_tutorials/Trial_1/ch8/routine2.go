package main

import (
	"fmt"
)

var naturals chan int
var square chan int

func main() {
	naturals = make(chan int)
	square = make(chan int)

	go counter(naturals)
	go squares(square, naturals)
	for n := range square {
		fmt.Printf("Data :%d\n", n)
	}
}

func counter(ch chan int) {
	for x := 0; x < 20; x++ {
		ch <- x
	}
	close(ch)
}

func squares(out, in chan int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}
