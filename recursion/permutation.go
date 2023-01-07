package main

import "fmt"

func permutate(a []rune, i int) {
	if i > len(a) {
		fmt.Println(string(a))
		return
	}

	permutate(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutate(a, i+1)
		a[i], a[j] = a[j], a[i]

	}
}

func main() {
	permutate([]rune("abc"), 0)
}
