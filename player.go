package main

import "fmt"

type Player struct {
	Name  string
	Cards []Card
	Call  string
	Out   bool
	Total int
}

func (p *Player) Print() {
	fmt.Printf("\t%s(%d):\n", p.Name, p.Total)
	for _, c := range p.Cards {
		c.Show()
	}
	fmt.Println()
}

func (p *Player) Take(cards []Card) {
	for _, c := range cards {
		p.Cards = append(p.Cards, c)
	}
}

func (p *Player) UpdateTotal() {
	ace := false
	p.Total = 0

	for _, c := range p.Cards {
		if c.RelValue == "Ace" {
			ace = true
		}
		p.Total = p.Total + c.Value
	}

	if ace && p.Total <= 11 {
		p.Total = p.Total + 10
	}
}
