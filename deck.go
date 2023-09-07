package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

// This is a receiver function
// print is the function name
// deck is a reference to the type that we want
// to attach this method to
// any variable of type deck it gets access now to the
// print method
// the variable d is an instance of the variable deck
// d is basically the receiver argument

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
