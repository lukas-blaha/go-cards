package main

func main() {
	deck := NewDeck()
	deck.Shuffle()

	players := []*Player{
		&Player{Name: "Lukas"},
		&Player{Name: "Tom"},
	}

	game := Game{
		Players: players,
		Dealer:  &Player{Name: "Dealer"},
		Deck:    &deck,
	}
	game.Play()
}
