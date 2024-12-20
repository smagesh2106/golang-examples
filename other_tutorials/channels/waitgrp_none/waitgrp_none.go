package main

import (
	"fmt"
)

func myfunc() {
	fmt.Println("Inside go routine")
}

func main() {
	fmt.Println("Hello world")
	go myfunc()
	fmt.Println("Finished Execution")
}
