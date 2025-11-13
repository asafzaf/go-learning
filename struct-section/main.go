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

	johnPointer := &p1

	johnPointer.updateName("Jane")
	johnPointer.print()

	p1.print()

	// Alternatively, i can call the method directly on the struct instance
	p1.updateName("Mike")
	p1.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
