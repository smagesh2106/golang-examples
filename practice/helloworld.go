package main

import (
	"fmt"
	"os"
	"strings"
)

func main1() {
	//fmt.Println("Hello world")
	var s, sep, tmp string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	tmp = ""
	j := 1
	for tmp != "five" && j < len(os.Args) {
		fmt.Println(tmp)
		tmp = os.Args[j]
		j++

	}
	fmt.Println(len(os.Args))
	fmt.Println(strings.Join(os.Args[1:], "-"))
}
