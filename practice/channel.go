package main

import (
	"fmt"
	"time"
)

func main27() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {

		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {

		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		time.Sleep(time.Second * 1)
		fmt.Println(x)
	}
}
