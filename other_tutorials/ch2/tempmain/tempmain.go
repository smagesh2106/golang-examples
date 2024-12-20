package main

import (
	t "examples/ch2/tempconv"
	"fmt"
)

func main() {
	fmt.Println(t.FToC(150))
	fmt.Println(t.CToF(50))
	t.SayHello("Paul..")
}
