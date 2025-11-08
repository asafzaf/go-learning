package main

import "fmt"

func main() {

	// First way to declare a variable
	var card string = "Ace of Spades"
	// Second way to declare a variable
	card2 := "Two of Hearts"
	// Re-assigning a variable
	card2 = "Three of Diamonds"

	var age int = 30
	var price float64 = 19.99
	var isAvailable bool = true
	isAvailable = false

	fmt.Println(card)
	fmt.Println(card2)
	fmt.Println(age)
	fmt.Println(price)
	fmt.Println(isAvailable)
}
