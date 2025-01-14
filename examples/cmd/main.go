package main

import (
	"fmt"

	"github.com/smagesh2106/examples/internal/greet"
	"github.com/smagesh2106/examples/internal/routes"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(greet.SayHello("Magesh"))
	fmt.Println(routes.Bye("Paul"))

}
