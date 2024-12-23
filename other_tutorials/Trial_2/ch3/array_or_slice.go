package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]interface{})
	m["a"] = []string{"a", "b", "c"}
	m["b"] = [4]int{1, 2, 3, 4}

	test(m)
}

func test(m map[string]interface{}) {
	for k, v := range m {
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Slice:
			fmt.Println(k, "is a slice with element type", rt.Elem())
		case reflect.Array:
			fmt.Println(k, "is an array with element type", rt.Elem())
		default:
			fmt.Println(k, "is something else entirely")
		}
	}
}
