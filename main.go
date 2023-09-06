package main

func main() {
	// card := "Ace of spades"
	// var card string = "Ace of spades" => this is the same as the above,
	// but we let Go compiler to dynamically assign the type
	// We only do this to initialise the value to the variable
	// When we want to change/assign the value of the variable
	// We only us = sign
	// card = "Five of Diamonds"
	// card := newCard()

	// Slice:
	/*cards := []string{
		newCard(), newCard(),
	}*/
	// Since we created the deck in the deck.go file, we can just call it
	// here instead of creating the whole slice
	cards := deck{
		newCard(), newCard(),
	}
	cards = append(cards, "Six of Spades")

	// We call the function from deck.go

	cards.print()

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
