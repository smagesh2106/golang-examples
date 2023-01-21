package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum4(arr))

}

func sum4(a []int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	} else {
		mid := len(a) / 2
		return (sum4(a[0:mid]) + sum4(a[mid:len(a)]))
	}
}
