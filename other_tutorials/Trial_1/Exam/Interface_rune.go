package main

import (
	"fmt"
	"strconv"
	//"strings"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (N *NumericInput) Add(x rune) {
	str := string(x)
	//	if strings.ContainsRune(str, x) {
	_, err := strconv.Atoi(str)
	if err == nil {
		N.input = N.input + string(x)
	}
	//	}
}

func (N *NumericInput) GetValue() string {
	return N.input
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
