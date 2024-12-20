package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func doSomething(i int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("-->:%d=%v\n", i, r.(string))
			err = errors.New(fmt.Sprintf("-->:%v", r.(string))) //<<<-- Best approch
		}
	}()

	tmp := rand.Intn(5)
	fmt.Printf("rand ->:%d\n", tmp)

	if tmp != 3 {
		panic("Panic")
	}
	return
}

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println("=================>> ", i)
		err := doSomething(i)
		if err != nil && i == 3 {
			for {
				time.Sleep(5 * time.Second)
				e := doSomething(i)
				if e == nil {
					break
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}
