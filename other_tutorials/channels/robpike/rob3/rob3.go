package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(s string) <-chan string {

	var c chan string
	c = make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}
func main() {
	//	var c chan string
	//	c = make(chan string)
	c1 := boring("Joe")
	c2 := boring("Ann")

	c := fanIn(c1, c2)
	for i := 0; i < 10; i++ {
		fmt.Printf("you say %q\n", <-c) //----<<<< both jan and ann are not blocked.
	}
	fmt.Println("You are boring, I am leaving...")
}
