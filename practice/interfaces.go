package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle2 struct {
	d float32
}

func (c Circle2) Area() float32 {
	return math.Pi * c.d * c.d
}

func (c Circle2) Perimeter() float32 {
	return 2 * math.Pi * c.d
}

type Square2 struct {
	s float32
}

func (s Square2) Area() float32 {
	return s.s * s.s
}

func (s Square2) Perimeter() float32 {
	return 4 * s.s
}

//================== Pointer type ===

// Define an interface
type Greeter interface {
	Greet() string
}

// Define a struct
type Person struct {
	Name string
}

// Define a method with a pointer receiver to implement the Greeter interface
func (p *Person) Greet() string {
	return "Hello, " + p.Name
}
func main15() {
	var shape Shape = Circle2{d: 3.5}
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
	shape = Square2{s: 4.0}
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	p := &Person{Name: "Alice"}
	var g Greeter = p
	fmt.Println(g.Greet())

	p2 := &Person{Name: "Mark"}
	var g2 Greeter = p2

	if _, ok := g2.(Greeter); ok {
		fmt.Println("type pointer to Greeter")
		fmt.Println(g2.Greet())
	}
}
