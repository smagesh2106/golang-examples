package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 3, 9, 8, 1, 5}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	fmt.Println(arr)
}
