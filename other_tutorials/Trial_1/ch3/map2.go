package main

import (
	"fmt"
)

type Employee struct {
	name string
	age  uint8
}

var m = make(map[string]string)

func main() {
	m["hello"] = "world-1"

	var emp Employee
	emp.name = "Magesh"
	emp.age = 51

	fmt.Println("Before map change...")
	fmt.Printf("%v\n", m)
	modify(m)
	fmt.Println("After map change...")
	fmt.Printf("%v\n", m)

	fmt.Println("\n\nBefore struct change...")
	fmt.Printf("%v\n", emp)
	modify2(&emp)
	fmt.Println("After struct change...")
	fmt.Printf("%v\n", emp)

	e1 := Employee{name: "foo", age: 1}
	e2 := Employee{name: "bar", age: 2}
	e3 := Employee{name: "foo", age: 1}
	fmt.Println(e1 == e2)
	fmt.Println(e1 == e3)
}

func modify(m map[string]string) {
	m["hello"] = "world-2"
}

func modify2(s *Employee) {
	s.name = "Paul"
}
