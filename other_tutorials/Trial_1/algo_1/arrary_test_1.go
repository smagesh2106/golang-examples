package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr[len(arr)-2], arr[len(arr)-1])
}
