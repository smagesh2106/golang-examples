package main

import "fmt"

func main() {
	a := 543222

	fmt.Println(add1(a))
}

func add1(n int) int {
	if n < 10 {
		return n
	} else {
		return add1(n/10) + n%10
	}

}
