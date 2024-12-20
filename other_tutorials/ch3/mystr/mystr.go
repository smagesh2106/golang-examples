package main

import (
	"fmt"
)

func main() {

	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(string(s[1]))
	fmt.Println(s[0:5])
	fmt.Println("Hello ....\U00004e16\U0000754c")
}
