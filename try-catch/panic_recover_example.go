package main

import (
	"fmt"
	"time"
)

func main() {
	loop1()
}

func loop1() {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case string:
				fmt.Println("string obj")
			case error:
				fmt.Println("error obj")
			default:
				fmt.Println("Unknown error")

			}
			fmt.Println("recovered from panic: ", r)
		}
	}()
	for i := 5; i >= 0; i-- {
		time.Sleep(500 * time.Millisecond)
		loop2(i)
	}
}

func loop2(i int) {
	fmt.Printf("------>In loop2 :%d\n", i)
	if i == 0 {
		//panic("encountered i==0")
		//panic(errors.New("encountered i==0"))
		panic(-1)
	}
}
