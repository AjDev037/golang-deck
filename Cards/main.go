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
	fmt.Println("------- Loading Hand ---------")
	hand := LoadDeck("my_hand", "hand")
	fmt.Println("------- Loading cards ---------")
	my_deck := LoadDeck("my_deck", "deck")
	fmt.Println("------- Deck loaded ---------")
	fmt.Println(my_deck.getDeckLength())
	if hand.getDeckLength() == 0 {
		fmt.Println("------- Getting hand ---------")
		hand, my_deck = getHand(my_deck)
	} else {
		fmt.Println("------- Your hand ---------")
		hand.printDeck()
	}

	fmt.Println(my_deck.getDeckLength())
	fmt.Println("------- Saving Hand ---------")
	hand.saveToFile("my_hand")
	fmt.Println("------- Saving Deck ---------")
	my_deck.saveToFile("my_deck")
}
