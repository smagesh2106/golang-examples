package main

import (
	"fmt"
)

func main() {
	f := Squares(0)
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}

}

func Squares(i int) func() int {
	var x int = i
	return func() int {
		x++
		return x * x
	}
}
