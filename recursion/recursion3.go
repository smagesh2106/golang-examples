package main

import (
	"fmt"
)

func moreFun(n int) {
	fmt.Println(n)
	if n > 2 {
		moreFun(n - 1)
		moreFun(n - 2)
		moreFun(n - 3)

	}

}
func main() {

	moreFun(5)
}
