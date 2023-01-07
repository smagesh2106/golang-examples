package main

import "fmt"

func permutate3(a []rune, l int, r int) {
	if l == r {
		fmt.Println(string(a))
		return
	}
	for i := l; i <= r; i++ {
		a[i], a[l] = a[l], a[i]
		permutate3(a, l+1, r)
		a[i], a[l] = a[l], a[i]
	}

}

func main() {
	a := []rune("abcde")
	permutate3(a, 0, len(a)-1)
}
