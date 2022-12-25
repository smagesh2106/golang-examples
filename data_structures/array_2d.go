package main

import (
	"fmt"
)

func main() {
	arr1 := [10][3]int{}
	var arr2 [5][]int
	arr3 := new([10][2]int)
	var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//var arr3 = new([5]int)
	//arr4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("\n\n")
	arr1[0][0] = 10
	arr1[0][1] = 20
	arr1[0][2] = 30
	arr3[0][0] = 4
	(*arr3)[0][1] = 8
	//arr2[0][0] = 1
	fmt.Println(arr1)
	fmt.Println(arr3)
	fmt.Println(*arr3)
	fmt.Println(arr4)

}
