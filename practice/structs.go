// structs.go
package main

import (
	"fmt"
	"math"
	"sync"
)

type Circle struct {
	diameter float32
}

func (c *Circle) area() float32 {
	return math.Pi * c.diameter * c.diameter
}

func (c *Circle) perimeter() float32 {
	return 2 * math.Pi * c.diameter
}

type Square struct {
	side float32
}

func (s *Square) area2() float32 {
	return s.side * s.side
}

func (s *Square) perimeter2() float32 {
	return 4 * s.side
}

type Container struct {
	Circle //annonymous
	Square //annonymous
	mu     sync.Mutex
}

func main14() {
	fmt.Println("Hello World!")
	var cont Container
	cont.Circle.diameter = 3.0
	cont.Square.side = 4.2
	cont.mu.Lock()
	defer cont.mu.Unlock()
	fmt.Println(cont.area())
	fmt.Println(cont.perimeter())
	fmt.Println(cont.area2())
	fmt.Println(cont.perimeter2())

}
