package main

import (
	"fmt"
)

func main() {
	var x complex128
	var y complex128

	x = complex(1, 2)
	y = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

}
