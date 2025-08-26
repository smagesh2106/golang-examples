package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		err := callRemote(i)
		if err != nil {
			fmt.Printf("Error on call %d: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded\n", i)
			break
		}
	}
}

func callRemote(i int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
			// or log the error, etc.
		}
	}()
	if i != 3 {
		time.Sleep(500 * time.Millisecond)
		panic("Simulated panic")
	} else {
		return nil // or return an error
	}
}
