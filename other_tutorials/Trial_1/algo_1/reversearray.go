package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr1)

	for i, j := 0, len(arr1)-1; i < len(arr1)/2; i, j = i+1, j-1 {
		arr1[i], arr1[j] = arr1[j], arr1[i]
	}
	fmt.Println(arr1)
}
