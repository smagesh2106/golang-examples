package main

import (
	"fmt"
)

func main() {
	arr1 := [10]int{}
	var arr2 [5]int
	var arr3 = new([5]int)
	arr4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr1[0] = 0
	arr1[1] = 1

	arr2[0] = 10
	arr2[1] = 20

	(*arr3)[0] = 100
	arr3[1] = 200

	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(*arr3)
	fmt.Println(arr4)
}
