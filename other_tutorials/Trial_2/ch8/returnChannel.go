package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(500 * time.Millisecond)
		}
		close(c) //close at the sending loc
	}()

	return c
}

func main() {
	c := boring("boring")
	for msg := range c {
		fmt.Printf("You say %q\n", msg)
	}

	fmt.Println("I am leaving you are boring.")
}
