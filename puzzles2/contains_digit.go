package main

import "fmt"

func main() {
	fmt.Println(contains_digit(899, 8))
}

func contains_digit(n int, d int) bool {
	if n < d {

		return n == d
	}
	if n%10 == d {
		return true
	} else {
		return contains_digit(n/10, d)
	}

}
