package main

import "fmt"

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
	cards := []string{
		newCard(), newCard(),
	}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
