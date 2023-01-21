package main

import "fmt"

func main() {
	fmt.Println(find_sum_2(5))
}

func find_sum_2(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n%2 != 0 {
		return (3*find_sum_2((n-1)/2) + find_sum_2(n+1)/2)
	} else {
		return (3*find_sum_2(n/2) + find_sum_2(n/2-1))
	}
}
