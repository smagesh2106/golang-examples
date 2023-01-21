package main

import "fmt"

func main() {
	fmt.Println(find_gcd(12, 8))
}

func find_gcd(m int, n int) int {
	if m == 0 {
		return n
	} else if m > n {
		return find_gcd(n, m)
	} else {
		return find_gcd(m, n-m)
	}
}
