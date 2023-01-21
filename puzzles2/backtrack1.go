package main

import "fmt"

func main() {
	arr := []rune{'a', 'b', 'c'}
	back1(arr, 0)
}

func back1(a []rune, n int) {
	if n == len(a) {
		fmt.Printf("++++++++++++++++++++> %v,  n  = %d\n", string(a), n)
		return
	}
	for i := n; i < len(a); i++ {
		fmt.Printf("----> %v, i = %d\n", string(a), i)
		a[i], a[j] = a[j], a[i]
		back1(a, i+1)
		a[i], a[j] = a[j], a[i]
		fmt.Printf("====> %v, i = %d\n", string(a), i)
	}
	return
}
