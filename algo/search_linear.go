package main

import (
	"fmt"
)

func linearSearch(items []int, key int) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{2, 34, 33, 36, 98, 47, 66, 22, 87, 88, 51, 48, 42, 59, 49, 92}
	fmt.Println(linearSearch(items, 49))
	fmt.Println(linearSearch(items, 46))
}
