package main

import "fmt"

var cache = [94]int64{}

func main() {
	print1(1)
	fmt.Println("================")
	print(1)
	fmt.Println("================")
	fmt.Println(cache)
	fmt.Println(fibb(92))
}
func print1(n int) {
	fmt.Println(n)
	print2(2)
}
func print2(n int) {
	fmt.Println(n)
	print3(3)
}
func print3(n int) {
	fmt.Println(n)
}

func print(n int) {
	if n == 3 {
		fmt.Println(n)
		return
	}
	fmt.Println(n)
	print(n + 1)
}

func fibb(n int64) int64 {
	if n == 2 || n == 1 {
		return 1
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = fibb(n-1) + fibb(n-2)
	return cache[n]

}
