package main

import "fmt"

type Card struct {
	Type     string
	RelValue string
	Value    int
}

var CardTypes = []string{"Hearts", "Diamods", "Clubs", "Spades"}
var CardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

func (c *Card) Show() {
	fmt.Printf("\t\t%s of %s\n", c.RelValue, c.Type)
}
