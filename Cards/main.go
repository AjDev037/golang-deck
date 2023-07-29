package main

import (
	"fmt"
)

func main() {

	//cards := newDeck()
	//hand := deck{}

	//cards.printDeck()
	//cards.getRandomCard()
	//fmt.Println("----------------")
	//I want to setup like, first it get the number of cards in the deck
	//then it get the first hand
	// displays the hand
	// and tell the remaining cards in the deck.
	//fmt.Println("----------------")
	//hand, cards = getHand(cards)
	//hand.printDeck()
	//fmt.Println("----------------")
	//cards.printDeck()
	//fmt.Println("----------------")
	//fmt.Println(cards.toString())
	//fmt.Println("------- Saving cards ---------")
	//cards.saveToFile("my_deck")
	fmt.Println("------- Loading cards ---------")
	deck := LoadDeck("my_deck")
	fmt.Println("------- Deck loaded ---------")
	fmt.Println(deck.getDeckLength())
}
