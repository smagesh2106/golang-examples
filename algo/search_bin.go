package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a = []int{1, 10, 20, 47, 59, 63, 74, 88, 99}
	//	var input = 100

	fmt.Println("Enter a number to search: ")
	var input string
	fmt.Scanln(&input)

	val, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Supplied value %s is not a number\n", input)
	} else {
		value := bin_search(a, val)
		value2 := bin2(a, val)
		value3 := bin3(&a, val)

		if value != -1 {
			fmt.Printf("Location-1 is :%d\n", value)
			fmt.Printf("Location-2 is :%d\n", value2)
			fmt.Printf("Location-3 is :%d\n", value3)
		} else {
			fmt.Println("Not found")
		}
	}
}

func bin_search(a []int, input int) int {

	start := 0
	end := len(a) - 1

	for start <= end {
		mid := start + ((end - start) / 2)

		if a[mid] == input {
			return mid
		}

		if input < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func bin2(a []int, num int) int {
	start := 0
	end := len(a) - 1

	for start <= end {
		mid := start + (end-start)/2

		if num == a[mid] {
			return mid
		}
		if num < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func bin3(b *[]int, num int) int {
	start := 0
	a := *b
	end := len(a) - 1

	for start <= end {
		mid := start + (end-start)/2

		if num == a[mid] {
			return mid
		}
		if num < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
