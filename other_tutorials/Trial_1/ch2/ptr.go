package main

import (
	"fmt"
)

type point struct {
	X int
	Y int
}

func main() {

	// var p type  =  expression
	var p point //expression is omited
	p.X = 10
	p.Y = 12

	var q = point{X: 10, Y: 12} //type is omited
	var r = point{X: 1, Y: 2}   //type is omited
	fmt.Printf("Before %v\n", p)
	add1(&p)
	fmt.Printf("After %v\n\n", p)

	fmt.Printf("Before %v\n", q)
	add1(&q)
	fmt.Printf("After %v\n", q)

	fmt.Printf("Address of q :%p\n", &q)
	fmt.Printf("Equals :%v\n", p == q)
	fmt.Printf("Equals :%v\n", p == r)
}

func add1(a *point) {
	tmp := point{X: 3, Y: 4}
	a.X = a.X + tmp.X //compiler prefixes '*' for you
	a.Y = a.Y + tmp.Y

}

func add2(a *point) {
	tmp := point{X: 3, Y: 4}
	(*a).X = (*a).X + tmp.X
	(*a).Y = (*a).Y + tmp.Y

}
