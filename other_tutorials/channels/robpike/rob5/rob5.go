package main

import (
	"fmt"
	//"math/rand"
	//"time"
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
			//time.Sleep(time.Duration(rand.Int63n(500) * time.Millisecond))
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

	//c := fanIn5(c1, c2)

	for i := 0; i < 15; i++ {

		select {

		case msg1 := <-c1:
			fmt.Printf("%v from c1\n", msg1.str)
			msg1.wait <- true

		case msg2 := <-c2:
			fmt.Printf("%v from c2\n", msg2.str)
			msg2.wait <- true
			/*
				case msg := <-c:
					fmt.Printf("%v from c\n", msg.str)
					msg.wait <- true
			*/
		}

		/*
			msg1 := <-c1
			fmt.Println(msg1.str)
			msg2 := <-c2
			fmt.Println(msg2.str)
			//msg1.wait <- true
			//msg2.wait <- true
		*/

	}
	fmt.Println("You are boring, I am leaving...")
}
