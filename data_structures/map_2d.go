package main

import "fmt"

func main() {
	m1 := map[string]map[string]int{
		"solomon": {"maths": 80, "english": 70},
		"mary":    {"maths": 74, "english": 90},
	}

	fmt.Println(m1)
}
