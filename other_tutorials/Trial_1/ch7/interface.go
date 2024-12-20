package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perm() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float32
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perm() float64 {
	return float64(2 * (r.height + r.width))
}

func (c circle) area() float64 {
	return float64(math.Pi * c.radius * c.radius)
}

func (c circle) perm() float64 {
	return float64(2 * math.Pi * c.radius)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perm())
}

func main() {
	r := rectangle{height: 4, width: 5}
	c := circle{radius: 5}

	fmt.Println("\nRectangle...")
	measure(r)
	fmt.Println("\n\nCircle...")
	measure(c)
}
