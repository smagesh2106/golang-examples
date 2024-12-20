package main

import (
	"fmt"
)

func main() {
	a := []int{1, 4, 7, 3, 12, 2, 10, 15, 5}
	swaped := true
	for j := 0; j < len(a)-1; j++ {

		if swaped {
			swaped = false
			for i := 0; i < len(a)-1; i++ {
				if a[i] > a[i+1] {
					a[i], a[i+1] = a[i+1], a[i]
					swaped = true
				}
			}
		} else {
			break
		}

	}
	fmt.Println(a)
}
