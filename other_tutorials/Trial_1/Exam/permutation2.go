package main

import (
	"fmt"
)

func generate(k int, arr []rune) {
	if k == 1 {
		fmt.Println(string(arr))
	} else {
		generate(k-1, arr)

		for i := 0; i < k-1; i += 1 {
			if k%2 == 0 {
				arr[i], arr[k-1] = arr[k-1], arr[i]
			} else {
				arr[0], arr[k-1] = arr[k-1], arr[0]
			}
			generate(k-1, arr)
		}
	}

}

func main() {
	arr := []rune("abc")
	generate(len(arr), arr)
}
