package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"foo1", "foo2"}
	b := []string{"foo3", "foo4"}
	p := [][]string{a, b}
	fmt.Println("p =", p)
	c := append(a, b...)

	fmt.Println(c)

	var str strings.Builder
	str.WriteString("hello")
	fmt.Println(str.String())
	str.WriteString(" world")
	fmt.Println(str.String())
	x := 1
	y := 2
	z := 3

	x, y, z = y, z, x

	fmt.Println(x, y, z)

}
