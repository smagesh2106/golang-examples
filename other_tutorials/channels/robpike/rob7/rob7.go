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

func boring(s string) <-chan Message {

	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", s, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			<-waitForIt
		}

	}()
	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s

			}
		}
	}()

	return c
}

func main() {
	c1 := boring("Joe")
	c2 := boring("Ann")

	c := fanIn(c1, c2)
	timeout := time.After(1 * time.Second)
	for {

		select {

		case msg := <-c:
			fmt.Printf("%v from c\n", msg.str)
			msg.wait <- true

		case <-timeout:
			fmt.Println("You talk too much...")
			return

		}

	}
	fmt.Println("You are boring, I am leaving...")
}
