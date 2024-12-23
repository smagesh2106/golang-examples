package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello world...")

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "--"
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], "++"))
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "=="
	}
	fmt.Println(s)

}
