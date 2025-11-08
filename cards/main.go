package main

import "fmt"

func main() {

	// First way to declare a variable
	// var card string = "Ace of Spades"
	// // Second way to declare a variable
	// card2 := "Two of Hearts"
	// // Re-assigning a variable
	// card2 = "Three of Diamonds"

	// var age int = 30
	// var price float64 = 19.99
	// var isAvailable bool = true
	// isAvailable = false

	// card3 := drawCard()
	// fmt.Println(card3)

	// fmt.Println(card)
	// fmt.Println(card2)
	// fmt.Println(age)
	// fmt.Println(price)
	// fmt.Println(isAvailable)

	cardsArray := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds", drawCard()}

	cardsArray = append(cardsArray, "Four of Clubs")

	fmt.Println(cardsArray)

	for i, card := range cardsArray {
		fmt.Println(i, card)
	}

	printState()
}

func drawCard() string {
	return "King of Clubs"
}
