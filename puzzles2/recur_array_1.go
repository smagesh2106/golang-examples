package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(find_arr_sum_1(arr))
}

func find_arr_sum_1(a []int) int {
	fmt.Println(a)
	if len(a) == 0 {
		return 0
	}
	return (a[0] + find_arr_sum_1(a[1:len(a)]))
}
