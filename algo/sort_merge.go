package main

import (
	"fmt"
)

func main() {
	//arr := []int{ 10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	arr := []int{10, 6, 2, 1, 5, 8}
	fmt.Println(arr)
	fmt.Println(mergeSort(arr))

}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	fmt.Printf("1:%v\t2:%v\n\n", first, second)
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i, j := 0, 0

	for (i < len(a)) && (j < len(b)) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
