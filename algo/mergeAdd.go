package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}

	fmt.Println(mergeAdd(arr))

}

func mergeAdd(arr []int) int {
	if len(arr) < 2 {
		return arr[0]
	}

	first := mergeAdd(arr[:len(arr)/2])
	second := mergeAdd(arr[len(arr)/2:])

	return arrayAddition(first, second)
}

func arrayAddition(a int, b int) int {
	return a + b
}
