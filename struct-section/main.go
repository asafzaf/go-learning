package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
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
	}
	p1.updateName("Jane")
	p1.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
