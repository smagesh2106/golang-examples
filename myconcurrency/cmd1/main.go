package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	go func() {
		data++
	}()
	time.Sleep(time.Second * 2)
	if data == 0 {
		fmt.Printf("the value  is 0\n")
	} else {
		fmt.Printf("the value -2- is %v.\n", data)
	}
}

//DOES NOT WORK
