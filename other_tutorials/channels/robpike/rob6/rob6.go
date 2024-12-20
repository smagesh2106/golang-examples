package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Message struct {
	str string
}

func boring(s string) <-chan Message {

	c := make(chan Message)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", s, i)}
			//time.Sleep(10 * time.Millisecond)
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

	//	c := fanIn(c1, c2)

	for i := 0; i < 15000; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("%v from c1\n", msg1.str)

		case msg2 := <-c2:
			fmt.Printf("%v from c2\n", msg2.str)
			/*
				case msg := <-c:
					fmt.Printf("%v from c\n", msg.str)
			*/
		}

		/*
			msg1 := <-c1
			fmt.Println(msg1.str)
			msg2 := <-c2
			fmt.Println(msg2.str)
		*/
		//msg1.wait <- true
		//msg2.wait <- true

	}
	fmt.Println("You are boring, I am leaving...")
}
