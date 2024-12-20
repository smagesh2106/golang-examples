package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) display() {
	p.X += 2
	p.Y += 2
	fmt.Printf("Inside X :%d, Y :%d\n", p.X, p.Y)
}

func (p *Point) displayPtr() {
	(*p).X += 2
	(*p).Y += 2
	fmt.Printf("Inside X :%d, Y :%d\n", (*p).X, (*p).Y)
}

func (p *Point) displayPtr2() {
	p.X += 2 //<--- auto de reference
	p.Y += 2
	fmt.Printf("Inside X :%d, Y :%d\n", p.X, p.Y)
}

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(float64(q.X-p.X), float64(q.Y-p.Y))
}

func main() {
	p := Point{X: 10, Y: 12}
	fmt.Printf("Before X :%d, Y :%d\n", p.X, p.Y)
	p.display()
	fmt.Printf("After  X :%d, Y :%d\n\n", p.X, p.Y)

	q := Point{X: 10, Y: 12}
	p1 := &q
	fmt.Printf("Before X :%d, Y :%d\n", (*p1).X, (*p1).Y)
	p1.displayPtr()
	fmt.Printf("After  X :%d, Y :%d\n\n", (*p1).X, (*p1).Y)

	q1 := Point{X: 10, Y: 12}
	p2 := &q1
	fmt.Printf("Before X :%d, Y :%d\n", p2.X, p2.Y)
	p2.displayPtr2()
	fmt.Printf("After  X :%d, Y :%d\n\n", p2.X, p2.Y)

	a := &Point{X: 8, Y: 10}
	b := &Point{X: 12, Y: 14}
	fmt.Printf("Distance :%f\n", a.Distance(b))

}
