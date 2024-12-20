package main

import (
	"fmt"
)

func main() {

	var s = []string{"jnpr", "emc", "ibm", "hpe"}

	for _, val := range s {
		fmt.Println(val)
		for i := 0; i < len(val); i++ {
			fmt.Printf("%c ", val[i])
		}
		fmt.Println()
	}
}
