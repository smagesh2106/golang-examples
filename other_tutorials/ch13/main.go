package main

import (
	"fmt"
	"sync"
	"time"
)

type Ball struct {
	hits int
}

var wg sync.WaitGroup

func main1() {
	table := make(chan *Ball, 3)
	quit1 := make(chan bool)
	quit2 := make(chan bool)
	go player("ping", table, quit1, &wg)
	go player("pong", table, quit2, &wg)
	wg.Add(2)
	//go player("ping", table)
	//go player("pong", table)
	table <- new(Ball)
	time.Sleep(time.Second * 1)
	<-table //game over
	quit1 <- true
	quit2 <- true
	wg.Wait()
	close(table)
	close(quit1)
	close(quit2)

	panic("show me stacks")

}

func player(name string, table chan *Ball, quit chan bool, wg *sync.WaitGroup) {
	//func player(name string, table chan *Ball) {
	/*
		for {
			ball := <-table
			ball.hits++
			fmt.Println(name, ball.hits)
			time.Sleep(time.Millisecond * 200)
			table <- ball
		}
	*/
label:
	for {
		select {
		case ball := <-table:

			ball.hits++
			fmt.Println(name, ball.hits)
			time.Sleep(time.Millisecond * 20)
			table <- ball
		case <-quit:
			fmt.Println(name, "I am quiting...")
			break label
		}
	}
	wg.Done()
	return
}

func f(left, right chan int) {
	left <- 1 + <-right

}
func main() {
	const n = 100000
	leftmost := make(chan int)
	left := leftmost
	//	right := leftmost
	//left := make(chan int)
	var right chan int

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
	//fmt.Println(<-left)
}
