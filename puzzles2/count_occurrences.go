package main

import "fmt"

func main() {
	arr := []int{1, 2, 7, 4, 4, 5, 6, 7, 7, 7, 8}
	fmt.Println(find_occur(arr, 7))
	fmt.Println(find_occur2(arr, 7))
}

func find_occur(a []int, x int) int {
	if len(a) == 0 {
		return 0
	}
	tmp := 0
	if a[0] == x {
		tmp = 1
	}
	return tmp + find_occur(a[1:], x)

}

func find_occur2(a []int, x int) int {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		if a[0] == x {
			return 1
		} else {
			return 0
		}
	} else {
		mid := len(a) / 2
		tmp := 0
		if a[0] == x {
			tmp = 1
		}

		b := find_occur2(a[1:mid], x)
		c := find_occur2(a[mid:len(a)], x)
		return tmp + b + c

	}
}
