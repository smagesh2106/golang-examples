package main

import (
	"fmt"
	"math"
)

//interface
type Geometry interface {
	area() float64
	perim() float64
}

//shapes
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//rectangle interface methods
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

//circle interface methods

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

/*
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
*/
//interface method
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	r := rect{width: 5, height: 4}
	c := circle{radius: 10}

	if v, ok := interface{}(r).(Geometry); ok {
		measure(v)
	}

	fmt.Println("\n\n")

	if v, ok := interface{}(c).(Geometry); ok {
		measure(v)
		fmt.Println("Circle OK")
	} else {
		fmt.Println("Circle NOT OK")
	}

	switch x := interface{}(r).(type) {
	case nil:
		fmt.Println("NIL")
	case rect:
		fmt.Printf("RECT :%v\n", x)

	}

}
