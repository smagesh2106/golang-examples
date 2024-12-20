package main

import (
	t "examples/ch2/tempconv"
	"fmt"
)

func main() {
	var c t.Celcius = 45
	fmt.Println("\n" + c.String())
}
