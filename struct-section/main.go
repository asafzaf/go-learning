package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		contact: contactInfo{
			email:   "john.doe@example.com",
			zipCode: 12345,
		},
		contactInfo: contactInfo{
			email:   "john.doe222@example.com",
			zipCode: 222222,
		},
	}

	p1.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
