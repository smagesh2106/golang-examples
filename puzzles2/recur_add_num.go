package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(find_sum(5, 0))
	end := time.Since(start)
	fmt.Println("Duration :", end.Nanoseconds())
}

func find_sum(n int, result int) int {
	result += n

	if n == 0 {
		return result
	}

	return find_sum(n-1, result)
}
