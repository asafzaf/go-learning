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
			email:   "john.doe@example.com",
			zipCode: 12345,
		},
	}
	fmt.Println(p1)
	p1.firstName = "Jane"
	p1.lastName = "Smith"
	p1.age = 28
	fmt.Println(p1)
	p1.age++
	fmt.Println(p1)
	fmt.Printf("%+v", p1)
}
