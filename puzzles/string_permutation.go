package main

import "fmt"

// Perm calls f with each permutation of a.

// Permute the values at index i to len(a)-1.

func perm(a []rune, i int) {
	if i > len(a) {
		fmt.Println(string(a))
		return
	}
	perm(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	perm([]rune("abc"), 0)
}
