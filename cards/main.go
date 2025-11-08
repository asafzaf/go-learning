package main

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

	// cardsArray := deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds", drawCard()}

	// cardsArray = append(cardsArray, "Four of Clubs")

	// cardsArray := newDeck()

	// hand, remainingDeck := deal(cardsArray, 5)

	// // cardsArray.print()

	// fmt.Println("Hand:")
	// hand.print()
	// fmt.Println("\nRemaining Deck:")
	// remainingDeck.print()

	// greeting := "Hello, welcome to the card game!"
	// fmt.Println([]byte(greeting))

	// printState()

	cards := newDeck()
	// fmt.Println(cards.toString()
	// err := cards.saveToFile("my_cards")
	// if err != nil {
	// 	fmt.Println("Error saving to file:", err)
	// }
	// loadedCards := newDeckFromFile("my_cards")
	// loadedCards.print()

	cards.shuffle()
	cards.print()
}

func drawCard() string {
	return "King of Clubs"
}
