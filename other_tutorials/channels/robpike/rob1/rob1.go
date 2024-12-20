package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(s string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", s, i)
		time.Sleep(time.Duration(rand.Int63n(3)) * time.Millisecond)
	}
}
func main() {
	var c chan string
	c = make(chan string)
	go boring("Joe", c)
	//time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("you say %q\n", <-c)
	}
	close(c)
	fmt.Println("You are boring, I am leaving...")
}
