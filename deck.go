package main

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	// Initialize a new deck of cards
	var d Deck
	var c Card

	for _, t := range CardTypes {
		for i, v := range CardValues {
			// for each type and value generate the card
			var ni int
			if i > 9 {
				ni = 10
			} else {
				ni = i + 1
			}

			c.Type = t
			c.RelValue = v
			c.Value = ni

			// append generated card to deck
			d.Cards = append(d.Cards, c)
		}
	}

	return d
}

func (d *Deck) Print() {
	for _, c := range d.Cards {
		c.Show()
	}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Hit(n int) []Card {
	var rc []Card
	rc = append(rc, d.Cards[:n]...)
	d.Cards = append(d.Cards[n:n+1], d.Cards[n+1:]...)
	return rc
}
