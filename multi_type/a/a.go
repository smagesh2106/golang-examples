package a

import "fmt"

type AWS struct {
	Name string
}

type AWSInterface interface {
	GetName()
}

func (a *AWS) GetName() {
	fmt.Println("AWS GetName: ", a.Name)
}

func NewAWS() *AWS {
	return &AWS{
		Name: "AWS Interface",
	}
}
