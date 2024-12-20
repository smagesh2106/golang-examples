package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}

func main() {

	var w Wheel
	w = Wheel{Circle: Circle{Center: Point{X: 1, Y: 2}, Radius: 10}, Spokes: 21}
	fmt.Println(w)

	var m Movie
	m.Title = "Matrix"
	m.Year = 1995
	m.Color = true
	m.Actors = []string{"Keanu Reves", "Morpheus"}

	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", data)
}
