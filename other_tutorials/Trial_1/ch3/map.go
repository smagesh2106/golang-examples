package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr"}
	var strs []string

	fmt.Println(months[2])
	for _, m := range months {
		strs = append(strs, m)

		fmt.Printf("%p\n", strs)
	}

	fmt.Println(strs)

}
