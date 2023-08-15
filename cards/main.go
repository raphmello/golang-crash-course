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

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
	printState()

}

func newCard() string {
	return "Five of Diamonds"
}
