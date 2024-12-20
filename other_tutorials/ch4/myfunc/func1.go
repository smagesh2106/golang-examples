package main

import (
	"fmt"
)

func main() {
	f := squares(2)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(divByTwo(10))
	defer func() {
		fmt.Printf("---> ")
		if p := recover(); p != nil {
			fmt.Println("cannot divide by zero")
		}
	}()
	for i := 10; i > -1; i-- {
		fmt.Println(divByTwo(0))
	}

}

func squares(x int) func() int {
	//var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func divByTwo(x int) float32 {
	if x == 0 {
		panic("input cannot be zero")
	} else {
		return float32(1 / x)
	}
}
