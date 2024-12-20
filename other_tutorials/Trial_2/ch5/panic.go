package main

import (
	"fmt"
	"time"
)

var count uint8

func foo() {
	count++
	if count < 5 {
		fmt.Printf("Count :%d\n", count)
		time.Sleep(3 * time.Second)
		panic("my exception\n")
	} else {
		fmt.Println("Don't panic")
	}
}
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("caught exception: %s", "Test")
			//foo()
			fmt.Println("......100")
		}
	}()
	fmt.Println("......1")
	foo()
	fmt.Println("......2")

}
