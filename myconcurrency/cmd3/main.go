package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	m    sync.Mutex
	data int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.m.Lock()
		defer v1.m.Unlock()

		time.Sleep(time.Second * 2)

		v2.m.Lock()
		defer v2.m.Unlock()
		fmt.Printf("Sum = %v \n", v1.data+v2.data)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

// DOES NOT work
