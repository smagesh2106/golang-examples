package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int)
	m2 := new(map[string]int)
	m3 := map[string]int{
		"foo": 1,
		"bar": 2,
		"haa": 3,
	}
	m1["apple"] = 2
	m1["orange"] = 3
	m1["banana"] = 4

	m2 = &m1

	fmt.Println(m1)
	fmt.Println(*m2)
	fmt.Println(m3)
}
