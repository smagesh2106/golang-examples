package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	done := make(chan interface{})

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		done <- ""
	}()

	//for i := 0; i < 2; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		case <-done:
			fmt.Println("recd DONE")
			os.Exit(0)
		}
	}
}
