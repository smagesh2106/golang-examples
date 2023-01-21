package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum3(arr))
}

func sum3(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return sum3(a[0:len(a)-1]) + a[len(a)-1]
}
