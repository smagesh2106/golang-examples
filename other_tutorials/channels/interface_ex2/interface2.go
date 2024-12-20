package main

import (
	"fmt"
)

type User struct {
	Firstname, Lastname string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

type Customer struct {
	Fullname string
}

func (c *Customer) Name() string {
	return fmt.Sprintf("%s", c.Fullname)
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main1() {
	u := &User{"Magesh", "Kumar"}
	fmt.Println(Greet(u))
	c := &Customer{"Paul-Amal-Sunder-Raj"}
	fmt.Println(Greet(c))

}
