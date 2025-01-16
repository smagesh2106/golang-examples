package main

import (
	"fmt"
	"os"
	"time"
)

func main26() {
	abort := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		abort <- struct{}{}
	}()
	for {
		select {

		case <-abort:
			fmt.Println("Aborted")
			os.Exit(0)
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Default")
		}

	}
	fmt.Println("hello")

}
