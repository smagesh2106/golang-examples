package main

import "fmt"

func main() {
	fmt.Println(fib2(10))
}

func fib2(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		aux := 1
		for i := 0; i < n-1; i++ {
			aux += fib2(i)
		}
		return aux
	}

}
