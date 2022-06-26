package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	for _, s := range cardSuits {
		for _, v := range cardValues {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

func (d deck) print() {
	for i, c := range d {
		fmt.Printf("%d-%s\n", i, c)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	bslice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error (%s) when loading file (%s) \n", err, filename)
		panic(err)
	}
	s := strings.Split(string(bslice), ",")
	d := deck(s)
	return d
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	upperLimit := len(d) - 1
	for i := range d {
		newPosition := r.Intn(upperLimit)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
