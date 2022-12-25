package main

import (
	"fmt"
)

func main() {
	count(10)

}

func count(n int) {
	if n == 1 {
		fmt.Println(n)
		return
	}
	fmt.Println(n)
	count(n - 1)
}
