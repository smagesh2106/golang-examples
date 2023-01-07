package main

import "fmt"

func permutate2(a []rune, l int) {
	if l == len(a) {
		fmt.Println(string(a))
		return
	}
	for i := l; i < len(a); i++ {
		a[i], a[l] = a[l], a[i]
		permutate2(a, l+1)
		a[i], a[l] = a[l], a[i]
	}

}

func main() {
	//a := []rune("abc")
	permutate2([]rune("abc"), 0)
}
