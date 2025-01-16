package main

import "fmt"

type Point struct {
	x int
	y int
	b bool
}

func main13() {
	var p Point
	p.x = 3
	p.y = 4

	//q.x = 1
	//q.y = 1
	q := Point{x: 1, y: 2, b: true}
	fmt.Println(p)
	(&p).update(&q)
	fmt.Println(p)
}

func (p *Point) update(q *Point) {
	p.x += q.x
	p.y += q.y
	p.b = q.b
}
