package main

import (
	"fmt"
	"time"
)

func main() {
	chanOwner := func() <-chan int {
		resStream := make(chan int)

		go func() {
			defer close(resStream)
			for i := 0; i < 5; i++ {
				resStream <- i
			}
		}()
		return resStream
	}

	dataStream := chanOwner()

	for dat := range dataStream {
		fmt.Printf("Recd %d\n", dat)
	}

	fmt.Println("Done Receiving")

	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

	var c <-chan int
	//close(c)
l1:
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Timed out.")
			break l1
		case <-c:
			time.Sleep(time.Second * 1)
			fmt.Println("Reading from a closed ch")
		}
	}

	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
