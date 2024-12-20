package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const N = 5
	var values [N]int32

	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		i := i
		log.Println("i -", i)
		go func() {
			log.Println("i =", i)
			values[i] = 50 + rand.Int31n(50)
			log.Println("Done :", i)
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("Values :", values)

	var ww sync.WaitGroup
	ww.Add(10)
	var a int = 10
	for i := 0; i < a; i++ {
		j := i
		go func() {
			log.Println("III -> ", j)
			ww.Done()
		}()

	}
	ww.Wait()

}
