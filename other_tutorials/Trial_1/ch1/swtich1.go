package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr = []int{1, 2, 3}

	for _, v := range arr {
		fmt.Printf("%d : ", v)
		switch getRandom() {
		case "heads":
			fmt.Println("its heads")
		case "tails":
			fmt.Println("its tails")
		default:
			fmt.Println("Landed on the edge")
		}
		time.Sleep(3 * time.Second)
	}
}

func getRandom() string {
	var n = rand.Uint64()
	fmt.Printf("Random number :%d\n", n)
	if n%2 == 0 {
		return "heads"
	} else {
		return "tails"
	}
}
