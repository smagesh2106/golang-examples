package main

import (
	"fmt"
)

func init() {
	fmt.Println("...init - 1")
}

func init() {
	fmt.Println("...init - 2")
}

func init() {
	fmt.Println("...init - 3")
}

type Greet = func(string)

func main() {
	fmt.Println("....main...")
	f := greet
	f("dog")
	fmt.Println("=============")
	func1(f)

	//function pointers
}

func func1(f Greet) {
	str := "hello"
	l := len(str)
	fmt.Println(l)
	sym := '-'
	for i, c := range str {
		if i == l-1 {
			sym = ' '
		}
		fmt.Printf("%c%c", c, sym)
	}
	fmt.Println()
	f("cat")
	fmt.Println("----------")
}

func greet(s string) {
	fmt.Println("hello world :" + s)
}
