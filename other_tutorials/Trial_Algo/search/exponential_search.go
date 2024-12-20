package main

import (
	"fmt"
)

func exponentialSearch(items []int, key int) bool {
	if items[0] == key {
		return true
	}

	i := 1
	for (i < len(items)) && (items[i] <= key) {
		i = i * 2
	}
	return binarySearch(items, i/2, key)

}

func binarySearch(items []int, low1 int, key int) bool {
	low := low1
	high := len(items) - 1

	if items[low] > key || items[high] < key {
		return false
	}

	if low == high && items[low] == key {
		return true
	}

	for low <= high {
		median := (low + high) / 2

		if items[median] == key {
			return true
		}

		if items[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return false
}

func main() {
	items := []int{2, 7, 9, 12, 14, 17, 21, 25, 29, 30, 33, 35, 38, 42, 44, 46, 49, 50, 53, 57, 63, 66, 69, 70}
	//items := []int{2}
	fmt.Println(exponentialSearch(items, 21))
}
