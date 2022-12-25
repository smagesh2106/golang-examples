package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 2)
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go swap(i, i+1)
		wg.Wait()
		fmt.Println(<-ch, <-ch)
	}
}

func swap(a, b int) {
	ch <- b
	ch <- a
	wg.Done()
}
