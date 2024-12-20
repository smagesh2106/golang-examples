package main

import (
	"errors"
	"fmt"
)

func test(i int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("++++>>>> Panic :%v", r.(string)))
		}
	}()
	if i == 3 {
		panic(fmt.Sprintf("Awww :%v", i))

	}
	return
}

func retryWrapper(i, try int) {
	err := test(i)
	fmt.Printf("i: %d try: %d\n", i, try)

	if err != nil && try < 5 {
		fmt.Println(err.Error())
		retryWrapper(i, try+1)
	}
}

func main() {
	for i := 1; i < 5; i++ {
		retryWrapper(i, 1)
	}
}
