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

	var c chan Message
	c = make(chan Message)
	fmt.Printf("Ref c :%v\n", c)

	waitForIt := make(chan bool)
	fmt.Printf("Ref waitForIt :%v\n", waitForIt)

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

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true // c1 and c2 are read sequencially
		msg2.wait <- true
	}
	fmt.Println("You are boring, I am leaving...")
}
