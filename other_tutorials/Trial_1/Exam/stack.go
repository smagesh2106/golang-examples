package main

import "fmt"

func main() {

	// Create

	var stack []string

	// Push

	stack = append(stack, "world!")

	stack = append(stack, "Hello ")
	stack = append(stack, "IBM ")

	for len(stack) > 0 {

		// Print top

		n := len(stack) - 1

		fmt.Print(stack[n])

		// Pop

		stack = stack[:n]

	}

	// Output: Hello world!

}
