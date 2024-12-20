package main

import (
	"fmt"
)

func main() {
	//func1()
	//	swap1()
	ptr()
}

func func1() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%f F = %g C\n", freezingF, fToC(freezingF))
	fmt.Printf("%f F = %g C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func swap1() {
	var i, j = 9, 10
	fmt.Printf("i=%v,  j=%v\n", i, j)
	i, j = j, i
	fmt.Printf("i=%v,  j=%v\n", i, j)
}

func ptr() {
	i := 1
	p := &i
	fmt.Printf("i=%v\n", i)
	*p = 2
	fmt.Printf("i=%v\n", i)
	fmt.Printf("i=%v\n", *p)

	iPtr := new(int)
	iPtr = p
	fmt.Printf("i=%v\n", *iPtr)
}
