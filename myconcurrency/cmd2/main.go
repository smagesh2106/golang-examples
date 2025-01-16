package main

import (
	"fmt"
	"sync"
)

var data int
var memLock sync.Mutex

func main() {

	go func() {
		memLock.Lock()
		data++
		memLock.Unlock()
	}()
	//time.Sleep(time.Second * 2)
	memLock.Lock()
	if data == 0 {
		fmt.Printf("the value  is 0\n")
	} else {
		fmt.Printf("the value -2- is %v.\n", data)
	}
	memLock.Unlock()
}

//DOES NOT WORK
