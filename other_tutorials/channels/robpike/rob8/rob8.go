package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring(s string, quit chan bool) <-chan Message {

	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- Message{fmt.Sprintf("%s %d", s, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			<-waitForIt
		}
		quit <- true
	}()
	return c
}
func cleanup() {
	fmt.Println("doing clean up....")
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}
func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)

	for {
		select {

		case msg := <-c:
			fmt.Printf("%v from c\n", msg.str)
			msg.wait <- true

		case <-quit:
			fmt.Println("Recd quit.....quitting...")
			cleanup()

			return

		}

	}
	fmt.Println("You are boring, I am leaving...")
}
