package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 3, 9, 8, 1, 5}

	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	fmt.Println(arr)
}
