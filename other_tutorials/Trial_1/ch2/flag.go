package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omit trailing newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	fmt.Printf("\nAddress of n : %p, value of n :%v\n", n, *n)
	fmt.Printf("Address of s : %p, value of s :%v\n", sep, *sep)

	intPtr := new(int)
	*intPtr = 5
	fmt.Printf("Address of prt : %p, value of prt : %v", intPtr, *intPtr)

}
