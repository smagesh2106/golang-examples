package main

import (
	"fmt"
	//"strconv"
)

const (
	Sunday int = iota
	Monday
	Tuesday
)

const (
	one uint = 1 << iota
	two
	three
	four
	five
	six
	seven
)

func main() {
	fmt.Println(HasPrefix("working in ibm", "work"))
	fmt.Println(HasSuffix("working in ibm", "ibm"))
	fmt.Println(contains("working in ibm", "is"))

	x := 123
	fmt.Println(string(x))
	fmt.Println(Tuesday)
	fmt.Printf("%b\n", seven)
}

func HasPrefix(s, prefix string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix
}

func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
