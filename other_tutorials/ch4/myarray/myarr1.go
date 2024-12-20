package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[2])

	b := [...]int{1, 2, 3}
	fmt.Println(b[1])

	symbol := [...]string{USD: "$", EUR: "\u20AC"}
	fmt.Println(USD, symbol[USD])
	fmt.Println(EUR, symbol[EUR])

	x := [2]int{1, 2}
	y := [...]int{1, 2}
	z := [2]int{1, 3}
	fmt.Println(x == y, x == z, y == z) // "true false  false"

	c1 := sha256.Sum256([]byte("Magesh"))
	fmt.Printf("%x\n", c1)
	ptrValues(&a)
	d := [5]int{1, 2, 3, 4, 5}
	reverse(d[:])

	fmt.Println(d)

	var runes []rune
	for _, r := range "Hello, World...." {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o'',' ' ' ' ' ' ']"

	var str string = "This is a string..."
	for _, ss := range str {
		fmt.Println(string(ss))
	}

}

func ptrValues(ptr *[3]int) {
	for i := range ptr {
		//ptr[i] = 0

		fmt.Printf("%d ", ptr[i])
		fmt.Printf("%d ", (*ptr)[i])
	}
}

func reverse(s []int) {
	for i, j := 0, (len(s) - 1); i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
