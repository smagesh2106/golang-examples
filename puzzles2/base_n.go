package main

import "fmt"

func main() {
	fmt.Println(find_base(15))
	x := find_base_iterative(15)
	for i := 0; i < len(x); i++ {
		fmt.Printf("%d", x[i])
	}

}

func find_base(n int) int {

	if n < 2 {
		return n
	}
	//r := n % 2
	//return 10*(find_base(n/5)) + n%5
	return 10*(find_base(n/2)) + n%2

}

func find_base_iterative(n int) [4]int {
	var x [4]int
	i := 0
	for n > 0 {
		x[i] = n % 2
		n = n / 2
		i++
	}
	return x
}
