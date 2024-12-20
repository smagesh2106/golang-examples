package main

import (
	"fmt"
	"sync"
)

func myfunc(w *sync.WaitGroup) {
	fmt.Println("Inside go routine")
	w.Done()
}

func main() {
	fmt.Println("Hello world")
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go myfunc(&waitGrp)

	go func() {
		fmt.Println("Inside my goroutine")
		waitGrp.Done()
	}()
	waitGrp.Wait()
	fmt.Println("Finished Execution")
}
