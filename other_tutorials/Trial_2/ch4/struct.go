package main

import (
	"fmt"
)

func main() {
	//data()
	//data2()
	//data3()
	//data4()

	data5()
}

type ArithOp func(int, int) int

func data5() {
	calculate(Plus)
	calculate(Minus)
	calculate(Multiply)
	calculate2(Plus)
	calculate2(Minus)
	calculate2(Multiply)

}

func calculate(fp func(int, int) int) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

// This is the same function but uses the type/fp defined above
//
func calculate2(fp ArithOp) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

//++++++++++++++++++++
type HelloFunc func(name string)

func data4() {

	type Emp struct {
		name  string
		age   uint8
		hello HelloFunc
	}

	e := Emp{
		name: "Paul",
		age:  51,
		//hello: func(n string) {
		//fmt.Printf("Hello %s, Welcome to the world\n", n)
		//},
		hello: SayHello,
	}
	e.hello(e.name)
}
func data3() {

	type Emp struct {
		name  string
		age   uint8
		greet HelloFunc
	}

	e1 := Emp{}
	e1.name = "Magesh"
	e1.age = 51
	e1.greet = SayHello
	e1.greet("Magesh")

}

func SayHello(name string) {
	fmt.Printf("Hello %s, Welcome to the world\n", name)
}

func data2() {
	m := make(map[string]uint8)
	m["Magesh"] = 51
	m["Paul"] = 52
	m["Amal"] = 53
	fmt.Println(m)

}
func data() {
	type Emp struct {
		name string
		age  uint8
	}

	var paul, amal Emp
	paul.name = "Paul"
	paul.age = 51

	amal.name = "Amal"
	amal.age = 52

	fmt.Println(paul)
	fmt.Println(amal)

	p := &paul
	p.age = 55
	fmt.Println(paul)
	pp := new(Emp)
	pp.name = "Magesh"
	pp.age = 60
	fmt.Println(*pp)
}
