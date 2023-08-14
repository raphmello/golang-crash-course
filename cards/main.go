package main

import "fmt"

var deckSize int

func main() {
	//declaring variables
	var card string = "Ace of Spades" //creating and assigning
	card1 := "Ace of Spades"          //creating and assigning
	card1 = "Five of Diamonds"        //only assigning
	fmt.Println(card)
	fmt.Println(card1)
	deckSize = 10
	fmt.Println(deckSize)
}
