package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{
		X: 2, Y: 3,
	}
	q := Point{
		X: 1, Y: 1,
	}

	fmt.Println(p.Add(q))
	fmt.Println(p.Add2(&q))
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Add2(q *Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}
