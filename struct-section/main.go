package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}
	fmt.Println(p1)
}
