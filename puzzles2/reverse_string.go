package main

import "fmt"

func main() {
	str := "hello"
	r := []rune(str)
	for i := len(r) - 1; i >= 0; i-- {
		fmt.Printf("%c", r[i])
	}
	fmt.Println("\n")
	fmt.Println(rev_string([]rune{'h', 'e', 'l', 'l', 'o'}))
}

func rev_string(s []rune) string {
	if string(s) == "" {
		return string("")
	}
	return rev_string(s[1:]) + string(s[0])
}
