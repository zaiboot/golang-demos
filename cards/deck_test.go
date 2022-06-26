package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	const LEN = 20
	if len(d) != LEN {
		t.Errorf("Length of deck should be %d, but got %d", LEN, len(d))
	}

	cards := map[int]string{0: "Ace of Spades", len(d) - 1: "Five of Clubs"}

	for i, c := range cards {
		ValidateCard(c, i, t, d)
	}

}
func ValidateCard(c1 string, i int, t *testing.T, d deck) {
	c := d[i]
	if c != c1 {
		t.Errorf("Expected %s, but got %s", c1, c)
	}
}
