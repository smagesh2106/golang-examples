package main

import (
	"fmt"
)

func MakeSquares(ch chan int, quit chan interface{}) {
	i := 1

	for {
		select {
		case ch <- (i * i):
			i++

		case <-quit:
			return
		}
	}
}

func main() {
	squares := make(chan int)
	quit := make(chan interface{})
	sum := 0
	go func() {
		for i := 0; i < 10; i++ {
			sum += <-squares
			fmt.Printf("---->%v\n", sum)
		}
		fmt.Println(sum)
		quit <- new(int)
	}()

	MakeSquares(squares, quit)
}
