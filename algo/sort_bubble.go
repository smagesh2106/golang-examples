package main

import (
	"fmt"
)

func main() {
	//var arr1 []int = []int{4, 3, 5, 8, 1, 9}
	var arr = []int{4, 3, 9, 8, 1, 5}
	fmt.Println(arr)
	for x := 0; x < len(arr)-1; x++ {
		for i := 0; i < len(arr)-x-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
