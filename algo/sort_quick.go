package main

import (
	"fmt"
	"math/rand"
)

//var arr = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 12, 16, 11, 17, 14, 13, 15, 20, 18, 19}

var arr = rand.Perm(20)

func main() {
	fmt.Printf("Before :%v\n", arr)
	quickSort(0, len(arr)-1)
	fmt.Printf("After :%v\n", arr)
}

func quickSort(low, high int) {
	if low < high {
		j := partition(low, high)
		quickSort(low, j)
		quickSort(j+1, high)
	}
}

func partition(low, high int) int {
	pivot := arr[low]
	i, j := low, high

	for i < j {

		for arr[i] <= pivot && i < high {
			i++
		}

		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}
