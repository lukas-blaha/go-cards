package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var mainCounter int

type Game struct {
	Players []*Player
	Dealer  *Player
	Deck    *Deck
	Out     int
	Rounds  string
	Last    bool
}

func (g *Game) Play() {
	// initial cards
	for _, p := range g.Players {
		p.Take(g.Deck.Hit(2))
		p.UpdateTotal()
	}
	g.Dealer.Take(g.Deck.Hit(2))

	g.printIntro()

	for {
		for _, p := range g.Players {
			callClear()
			g.printPlayers(false)
			if !p.Out {
				fmt.Printf("\t%s's call: ", p.Name)
				sc := bufio.NewScanner(os.Stdin)
				for sc.Scan() {
					if g.checkHit(sc.Text()) {
						n := g.parseHit(sc.Text())
						p.Take(g.Deck.Hit(n))
						p.UpdateTotal()
						p.Call = "hit"
						if p.Total > 21 {
							p.Out = true
							fmt.Printf("\n\t%s busted!\n", p.Name)
							p.Call = "stay"
							time.Sleep(2 * time.Second)
						}
						break
					} else if strings.ToLower(sc.Text()) == "stay" || strings.ToLower(sc.Text()) == "s" {
						fmt.Printf("\n\t%s calls to stand!\n", p.Name)
						p.Call = "stay"
						break
					} else {
						fmt.Printf("\tIncorrect call, you have two options: [hit/stay]\n")
						fmt.Printf("\t%s's call: ", p.Name)
					}
				}
			}
		}

		g.checkPlayers()

		g.checkCalls()
		if mainCounter > 1 {
			g.playDealer()
			callClear()
			g.printPlayers(true)
			g.final()
		}
	}
}

func (g *Game) final() {
	var winner Player
	var draw Player

	for _, p := range g.Players {
		if p.Total > winner.Total && p.Total <= 21 {
			winner = *p
		} else if p.Total == winner.Total {
			draw = *p
		}
	}

	if g.Dealer.Total > winner.Total && g.Dealer.Total < 21 {
		fmt.Printf("\n\tHouse wins the game!\n\n")
	} else if winner.Total == draw.Total && winner.Total == g.Dealer.Total {
		fmt.Printf("\n\tNo winners!\n\n")
	} else if winner.Total == g.Dealer.Total {
		fmt.Printf("\n\tIt's a draw between %s and house!\n\n", winner.Name)
	} else if winner.Total == draw.Total {
		fmt.Printf("\n\tIt's a draw between players %s and %s!\n\n", winner.Name, draw.Name)
	} else {
		fmt.Printf("\n\t%s wins the game!\n\n", winner.Name)
	}
	os.Exit(0)
}

func (g *Game) checkHit(s string) bool {
	re1 := regexp.MustCompile("^h$")
	re2 := regexp.MustCompile("^h [0-9]*$")
	re3 := regexp.MustCompile("^hit$")
	re4 := regexp.MustCompile("^hit [0-9]*$")

	if len(re1.FindAll([]byte(s), -1)) > 0 {
		return true
	} else if len(re2.FindAll([]byte(s), -1)) > 0 {
		return true
	} else if len(re3.FindAll([]byte(s), -1)) > 0 {
		return true
	} else if len(re4.FindAll([]byte(s), -1)) > 0 {
		return true
	}

	return false
}

func (g *Game) parseHit(s string) int {
	hit := strings.Split(s, " ")
	if len(hit) == 2 {
		n, _ := strconv.Atoi(hit[1])
		return n
	}

	return 1
}

func (g *Game) playDealer() {
	var ace bool
	d := g.Dealer
	d.UpdateTotal()

	for {
		callClear()
		g.printPlayers(true)

		for _, c := range d.Cards {
			if c.RelValue == "Ace" {
				ace = true
			} else {
				ace = false
			}
		}

		if d.Total > 16 {
			d.UpdateTotal()
			return
		} else if ace {
			if d.Total+10 <= 16 {
				d.Take(g.Deck.Hit(1))
				fmt.Printf("\tDealer has to hit!\n")
				time.Sleep(4 * time.Second)
			}
		} else if d.Total <= 16 {
			d.Take(g.Deck.Hit(1))
			fmt.Printf("\tDealer has to hit!\n")
			time.Sleep(4 * time.Second)
		}

		d.UpdateTotal()
	}
}

func (g *Game) checkPlayers() {
	var out int
	for _, p := range g.Players {
		if p.Out {
			out++
		}
	}

	if out == len(g.Players) {
		callClear()
		g.printPlayers(false)
		fmt.Printf("\n\tAll players busted! House wins!\n\n")
		os.Exit(0)
	}
}

func (g *Game) printPlayers(dealer bool) {
	for _, p := range g.Players {
		p.Print()
	}
	if dealer {
		g.Dealer.Print()
	}
}

func (g *Game) checkCalls() {
	var count int
	for _, p := range g.Players {
		if p.Call == "stay" {
			count++
		}
	}

	if count == len(g.Players) {
		mainCounter++
	} else {
		mainCounter = 0
	}

	if mainCounter == 1 {
		callClear()
		g.printPlayers(false)
		fmt.Printf("\tLast round if any of the players does not hit!\n")
		time.Sleep(3 * time.Second)
	}
}

func (g *Game) printIntro() {
	for i := 5; i > 0; i-- {
		callClear()
		fmt.Printf("\tWelcome to \"Go\" Blackjack\n")
		fmt.Printf("\n\tPlayer's options:\n")
		fmt.Printf("\t  - hit | h [number of cards]\n\t  - stay | s\n\n")

		fmt.Printf("\n\n\tGames begins in %d sec\n\n", i)
		time.Sleep(1 * time.Second)
	}

	callClear()
	fmt.Printf("\n\n\tEnjoy the game :)\n")
	time.Sleep(2 * time.Second)
}

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()
}
