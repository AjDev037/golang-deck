package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) getRandomCard() {
	value := rand.Intn(len(d))

	fmt.Println("Random card picked: ", d[value])
}

//func (d deck) getDeckLength() {
//	fmt.Println(len(d))
//}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

func getHand(d deck) (deck, deck) {

	hand := d[:5]
	d = d[5:]

	return hand, d
}

func (d deck) toString() string {
	stringsSlice := []string(d)
	valueString := strings.Join(stringsSlice, ",")

	return valueString
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
