package main

import (
	"fmt"
)

func main() {

	seen := make(map[string]struct{}) // set of strings
	seen["foo"] = struct{}{}
	seen["bar"] = struct{}{}

	e, k := seen["foo1"]
	fmt.Println(k, e)
	if _, ok := seen["foo1"]; ok {
		//seen[s] = struct{}{}
		fmt.Println("found")

	} else {
		fmt.Println("not found")
	}
}
