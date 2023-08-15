package main

//var deckSize int

//func main() {
//	//declaring variables
//	var card string = "Ace of Spades" //creating and assigning
//	card1 := "Ace of Spades"          //creating and assigning
//	card1 = "Five of Diamonds"        //only assigning
//	fmt.Println(card)
//	fmt.Println(card1)
//	deckSize = 10
//	fmt.Println(deckSize)
//}

//func main() {
//	card := newCard()
//	fmt.Println(card)
//
//}
//
//func newCard() string {
//	printState()
//	return "Five of Diamonds"
//}

// cards[0:2] returns a slice containing index 0 and 1
// cards[:2] returns a slice containing all indexes from beginning up to 1
// cards[1:] returns a slice containing index 1 up to the end
func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 4)

	hand.print()
	remainingCards.print()
}
