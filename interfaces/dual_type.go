package main

import "fmt"

func doTest(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int :", v)
	case float64:
		fmt.Println("float64 :", v.(float64)+2.1)
	case string:
		fmt.Println("string :", v.(string)+" world")
	default:
		fmt.Println("default")
	}
}

func main() {
	doTest(1)
	doTest(1.1)
	doTest("Hello")

}
