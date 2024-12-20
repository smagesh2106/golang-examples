package main

import (
	"fmt"
)

func main() {
	printNumbers(1, 2, 3, 4, 5)
	fmt.Println("\n\n")
	printNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func printNumbers(nos ...int) {
	for _, v := range nos {
		fmt.Println(v)
	}
}
