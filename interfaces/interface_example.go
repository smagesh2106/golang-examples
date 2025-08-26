package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type MultiStruct struct {
	Object   Shape
	Function func(int) int
}

func main() {
	var s Shape = Circle{Radius: 5.0}
	fmt.Printf("Area of a Circle :%.2f\n", s.Area())
	fmt.Printf("Perimeter of a Circle :%.2f\n", s.Perimeter())

	s = Rectangle{Length: 3.0, Width: 4.0}
	fmt.Printf("Area of a Rectable :%v\n", s.Area())
	fmt.Printf("Perimeter of a Rectangle :%v\n", s.Perimeter())

	var ms = MultiStruct{
		Object: &Circle{Radius: 5},
		Function: func(i int) int {
			return i * i
		},
	}
	obj1 := ms.Object
	fmt.Printf("----> Area :%v \n", obj1.Area())
	fmt.Printf("----> Perimeter :%v \n", ms.Object.Perimeter())
	fmt.Printf("----> Square :%v \n\n", ms.Function(6))

	ms = MultiStruct{
		Object: &Rectangle{Length: 10, Width: 8},
		Function: func(i int) int {
			return i / 2
		},
	}

	fmt.Printf("----> Area :%v \n", ms.Object.Area())
	fmt.Printf("----> Perimeter :%v \n", ms.Object.Perimeter())
	fmt.Printf("----> Half :%v \n", ms.Function(6))

}
