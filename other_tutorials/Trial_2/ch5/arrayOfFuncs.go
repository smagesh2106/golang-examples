package main

import (
	"fmt"
)

func main() {
	var sayHello []func()

	sayHello = append(sayHello, func() {
		fmt.Println("Hello ....1")
	})
	sayHello = append(sayHello, func() {
		fmt.Println("Hello ....2")
	})

	sayHello = append(sayHello, func() {
		fmt.Println("Hello ....2")
	})

	for _, f := range sayHello {
		f()
	}
}
