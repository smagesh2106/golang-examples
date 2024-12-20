package main

import (
	"fmt"
	"time"
)

func generate(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(5 * time.Millisecond)
	}
	close(c) //make sure to close the channel
}

func main() {
	ch := make(chan int, 5)

	go generate(ch)
	time.Sleep(10 * time.Millisecond)
	go generate(ch)
	time.Sleep(10 * time.Millisecond)

	go generate(ch)
	time.Sleep(10 * time.Millisecond)

	for v := range ch {
		fmt.Println(v)
	}
}
