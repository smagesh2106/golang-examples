package main

import (
	"fmt"
)

func main() {
	sl1 := [][]int{}
	sl2 := new([][]int)
	fmt.Println(sl1)
	fmt.Println("\n\n")
	sl1 = append(sl1, []int{1, 2, 3}, []int{1, 2, 3, 4})
	*sl2 = append(*sl2, []int{3, 3, 3}, []int{1, 2, 3, 4}, []int{4, 3, 2, 1})
	fmt.Println(sl1)
	fmt.Println(sl2)

}
