package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m    sync.Mutex
	data int
}

func (c *Counter) Incr() {
	c.m.Lock()
	defer c.m.Unlock()
	c.data++
}

func main() {
	var ctr Counter
	var wg sync.WaitGroup
	//ctr.data = 0

	wg.Add(1000000)

	for i := 0; i < 1000000; i = i + 1 {
		go func(ctr *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			ctr.Incr()
			//time.Sleep(time.Millisecond * 250)
		}(&ctr, &wg)
	}

	wg.Wait()
	fmt.Printf("Data = %v\n", ctr.data)

}

// Working
