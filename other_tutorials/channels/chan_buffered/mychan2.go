package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculateValue(c chan int) {
	value := rand.Intn(10)
	fmt.Printf("calculated Random value %d\n", value)
	time.Sleep(1000 & time.Millisecond)
	c <- value
	fmt.Println("Waiting for someone to receive the value from the channel")
}

func main() {
	fmt.Println("Go Channel Tutorial")
	//valueChan := make(chan int) //unbuffered
	valueChan := make(chan int, 2) //buffered
	defer close(valueChan)
	go calculateValue(valueChan)
	go calculateValue(valueChan)
	values := <-valueChan
	fmt.Println(values)
	values = <-valueChan
	fmt.Println(values)

}
