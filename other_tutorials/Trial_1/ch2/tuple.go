package main

import (
	"fmt"
)

func main() {
	fmt.Printf(" GCD of 5, 15 is %d\n", gcd(5, 15))
	fmt.Printf(" GCD of 6, 8 is %d\n", gcd(6, 8))

	fmt.Printf("6%8 = %v\n", 6%8)
	fib(9)
}

func gcd(x, y int) int {

	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) {

	x, y := 0, 1
	fmt.Printf("%d ", x)
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Printf("%d ", x)
	}
}
