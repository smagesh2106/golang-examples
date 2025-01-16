package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	//m    sync.Mutex
	data int
}

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		l.Lock()

		//v1.m.Lock()
		//defer v1.m.Unlock()

		time.Sleep(time.Second * 1)

		//v2.m.Lock()
		//defer v2.m.Unlock()
		fmt.Printf("Sum = %v \n", v1.data+v2.data)
		v1.data = 1
		v2.data = 1
		l.Unlock()
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}

// Works
