package main

import (
	"fmt"
)

func main() {
	var sl1 []int
	sl2 := make([]int, 5)
	sl := new([]int)
	sl3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)

	//sl1[0] = 1
	sl1 = append(sl1, 3)
	sl2[4] = 4
	sl3 = append(sl3, 3, 4, 5)
	fmt.Println("\n\n")
	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)
	fmt.Println("\n\n")
	i := 2
	sl3 = append(sl3[:i], sl3[i+2:]...)
	fmt.Println(sl3)
	fmt.Println(sl)
	*sl = append(*sl, 1, 2, 3, 4, 5)
	fmt.Println(sl)
	fmt.Println(*sl)
}
