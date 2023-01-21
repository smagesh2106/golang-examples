package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(find_arr_sum_2(arr, 0))
}

func find_arr_sum_2(a []int, sum int) int {
	if len(a) == 0 {
		return 0
	}

	sum += a[0] + find_arr_sum_2(a[1:len(a)], sum)
	return sum
}
