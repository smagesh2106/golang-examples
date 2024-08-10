package b

import "fmt"

type SRO struct {
	Name string
}

type SROInterface interface {
	GetName()
}

func (a *SRO) GetName() {
	fmt.Println("SRO GetName: ", a.Name)
}

func NewSRO() *SRO {
	return &SRO{
		Name: "SRO Interface",
	}
}
