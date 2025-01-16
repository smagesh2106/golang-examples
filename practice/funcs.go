package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main12() {
	fmt.Println(hypot(2.5, 3.5))

	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	label := true
	for label {
		err := retryOperation2()
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second)
		} else {
			label = false
			fmt.Println("Retry operation passed")
		}
	}
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func riskyOperation2() {
	r := rand.Intn(25)
	fmt.Println("---->", r)
	if r != 16 {
		panic(fmt.Sprintf("panic on attempt %d", r))
	} else {
		fmt.Printf("Operation successful on attempt \n")
	}
}

// Retry function with panic and recover logic
func retryOperation2() (err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic message and continue retrying
			fmt.Printf("Recovered from panic: %v. Retrying...\n", r)
			err = errors.New("panic")
		}

	}()

	// Try to perform the risky operation
	riskyOperation2()
	fmt.Println("Finished retrying.")
	return nil

}
