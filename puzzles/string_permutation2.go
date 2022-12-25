package main

import "fmt"

func perm2(a []rune, i int) {
	if i > len(a) {

		fmt.Printf("---------------------->0, i=%d,  a=%s\n", i, string(a))
		return
	}
	fmt.Printf("###############> i=%d,  a=%s\n", i, string(a))
	perm2(a, i+1)
	//fmt.Printf("--->1, i=%d, a=%s\n", i, string(a))

	for j := i + 1; j < len(a); j++ {
		fmt.Printf("===>2, i=%d,j=%d, a=%s\n", i, j, string(a))
		a[i], a[j] = a[j], a[i]
		fmt.Printf("===>3, i=%d,j=%d, a=%s\n", i, j, string(a))
		perm2(a, i+1)
		fmt.Printf("+++>4, i=%d,j=%d, a=%s\n", i, j, string(a))
		a[i], a[j] = a[j], a[i]
		fmt.Printf("+++>5, i=%d,j=%d, a=%s\n", i, j, string(a))
	}

}

func main() {
	perm2([]rune("abc"), 0)
}
