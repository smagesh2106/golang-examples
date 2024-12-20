package main

import (
	"fmt"
)

func interpolationSearch(items []int, key int) bool {
	low := 0
	high := len(items) - 1

	if items[low] > key || items[high] < key {
		return false
	}

	for low <= high && key >= items[low] && key <= items[high] {
		if low == high {
			if items[low] == key {
				return true
			}
			return false
		}

		mid := low + ((high-low)/(items[high]-items[low]))*(key-items[low])

		if items[mid] == key {
			return true
		} else {
			if items[mid] < key {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}

func main() {
	items := []int{2, 7, 9, 12, 14, 17, 21, 25, 29, 30, 33, 35, 38, 42, 44, 46, 49, 50, 53, 57, 63, 66, 69, 70}
	fmt.Println(interpolationSearch(items, 69))
}
