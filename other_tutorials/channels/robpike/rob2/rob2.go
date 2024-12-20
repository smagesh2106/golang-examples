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
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
func main() {
	//	var c chan string
	//	c = make(chan string)
	c1 := boring("Joe")
	c2 := boring("Ann")
	//time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Printf("you say %q\n", <-c1)
		fmt.Printf("you say %q\n", <-c2) //<<<<---c1 and c2 block each other
	}
	fmt.Println("You are boring, I am leaving...")
}
