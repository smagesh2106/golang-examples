package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main9() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	for s, age := range ages {
		fmt.Printf("%s\t%d\n", s, age)
	}

	age, ok := ages["alice"]
	if ok {
		fmt.Println(age)
	}

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

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	fmt.Println(w)

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

}
