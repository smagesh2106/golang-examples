package main

import (
	"examples/ch12/database"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello world...")
	var p1 database.Point
	var p2 database.Point
	var pts []database.Point
	//var pts = []database.Point{database.Point{X: 11, Y: 22}, database.Point{X: 111, Y: 222}}

	p1.X = 10
	p1.Y = 20

	p2.X = 100
	p2.Y = 200

	//pts = []database.Point{database.Point{X: 11, Y: 22}, database.Point{X: 111, Y: 222}}

	//pts[0] = p1
	//pts[1] = p2
	database.Change(&pts)
	fmt.Printf("%v", pts)

	var n1 string = ""
	var n2 string = "asd"
	var n3 int = 123
	var n4 float32 = 2.2

	fmt.Println("\n\n" + reflect.TypeOf(n1).Name())
	fmt.Println("\n\n" + reflect.TypeOf(n2).Name())
	fmt.Println("\n\n" + reflect.TypeOf(n3).Name())
	fmt.Println("\n\n" + reflect.TypeOf(n4).Name())
	if reflect.ValueOf(n1).Kind() == reflect.Int {
		//if reflect.TypeOf(n3).Name() == "int" {
		fmt.Println("n1 is numeric")
	} else {
		fmt.Println("n1 is not numeric")
	}

}
