package main

import (
	"fmt"
	"math/rand"
)

type Mana int

const (
	Any Mana = iota
	White
	Blue
	Black
	Red
	Green
)

type Type int

const (
	Creature Type = iota
	Instant
	Sorcery
	Enchantment
	Artifact
	Land
)

type Card struct {
	Type      Type
	Name      string
	Cost      []Mana
	Produce   []Mana
	Power     int
	Toughness int
}

func (c *Card) CanProduce(m Mana) bool {
	for _, p := range c.Produce {
		if p == m {
			return true
		}
	}
	return false
}

var ChiefOfTheEdge = &Card{
	Type:      Creature,
	Name:      "Chief of the Edge",
	Cost:      []Mana{White, Black},
	Power:     3,
	Toughness: 2,
}

var Swamp = &Card{
	Type:    Land,
	Name:    "Swamp",
	Produce: []Mana{Black},
}

var Plains = &Card{
	Type:    Land,
	Name:    "Plains",
	Produce: []Mana{White},
}

type Deck struct {
	Cards []*Cards
}

type Cards struct {
	Card   *Card
	Amount int
}

var MarduWorrier = &Deck{
	Cards: []*Cards{
		{
			Card:   ChiefOfTheEdge,
			Amount: 4,
		},
		{
			Card:   Swamp,
			Amount: 28,
		},
		{
			Card:   Plains,
			Amount: 28,
		},
	},
}

type Library []*Card

func (l Library) Shuffle() {
	for i := 0; i < len(l)-1; i++ {
		j := rand.Intn(len(l)-i) + i
		l[i], l[j] = l[j], l[i]
	}
}

func MakeLibrary(deck *Deck) Library {
	var l Library
	for _, cs := range deck.Cards {
		for i := 0; i < cs.Amount; i++ {
			l = append(l, cs.Card)
		}
	}
	return l
}

type CardInPlay struct {
	Tapped            bool
	SummoningSickness bool
	Card              *Card
	Game              *Game
}

func (c *CardInPlay) Power() int {
	return c.Card.Power
}

type Game struct {
	Turn         int
	Life         int
	OpponentLife int
	First        bool
	Hand         []*Card
	Library      []*Card
	BattleField  []*CardInPlay
}

func (g *Game) Print() {
	fmt.Printf("Turn: %d\n", g.Turn)
	fmt.Printf("Life: %d\n", g.Life)
	fmt.Printf("OpponentLife: %d\n", g.OpponentLife)
	fmt.Printf("First: %t\n", g.First)
	fmt.Printf("Hand (%d):\n", len(g.Hand))
	for i, c := range g.Hand {
		fmt.Printf("%d: %s\n", i, c.Name)
	}
	fmt.Printf("Library (%d):\n", len(g.Library))
	for i, c := range g.Library {
		fmt.Printf("%d: %s\n", i, c.Name)
	}
	fmt.Printf("BattleField (%d):\n", len(g.BattleField))
	for i, c := range g.BattleField {
		fmt.Printf("%d: %s\n", i, c.Card.Name)
	}
}

type Status int

const (
	Playing Status = iota
	Win
	Lose
	Draw
)

func (g *Game) PlayOneTurn() Status {
	g.Turn++
	g.Untap()
	if s := g.Draw(); s != Playing {
		return s
	}
	if s := g.FirstMain(); s != Playing {
		return s
	}
	if s := g.Combat(); s != Playing {
		return s
	}
	if s := g.SecondMain(); s != Playing {
		return s
	}
	g.Discard()
	return Playing
}

func (g *Game) Untap() {
	for _, c := range g.BattleField {
		c.Tapped = false
		c.SummoningSickness = false
	}
}

func (g *Game) Draw() Status {
	if g.First && g.Turn == 1 {
		return Playing
	}
	if len(g.Library) == 0 {
		return Lose
	}
	g.Hand = append(g.Hand, g.Library[0])
	g.Library = g.Library[1:]
	return Playing
}

func Take(c []*Card, i int) []*Card {
	if i == len(c) {
		return c[0:i]
	}
	return append(c[0:i], c[i+1:]...)
}

func (g *Game) FirstMain() Status {
	// Put a land.
	for i, c := range g.Hand {
		if c.Type == Land {
			g.BattleField = append(g.BattleField, &CardInPlay{
				Tapped:            false,
				SummoningSickness: false,
				Card:              c,
				Game:              g,
			})
			g.Hand = Take(g.Hand, i)
			break
		}
	}

	// Put a creature.
	for i, c := range g.Hand {
		if c.Type == Creature {
			j := 0
			var lands []*CardInPlay
			for _, cip := range g.BattleField {
				if j == len(c.Cost) {
					break
				}
				if !cip.Tapped && cip.Card.CanProduce(c.Cost[j]) {
					j++
					lands = append(lands, cip)
				}
			}
			if j == len(c.Cost) {
				for _, l := range lands {
					l.Tapped = true
				}
				g.BattleField = append(g.BattleField, &CardInPlay{
					Tapped:            false,
					SummoningSickness: true,
					Card:              c,
					Game:              g,
				})
				g.Hand = Take(g.Hand, i)
			}
		}
	}
	return Playing
}

func (g *Game) Combat() Status {
	for _, c := range g.BattleField {
		if c.Tapped || c.SummoningSickness {
			continue
		}
		c.Tapped = true
		g.OpponentLife -= c.Power()
	}
	if g.OpponentLife <= 0 {
		return Win
	}
	return Playing
}

func (g *Game) SecondMain() Status {
	return Playing
}

func (g *Game) Discard() {
	if len(g.Hand) > 7 {
		g.Hand = g.Hand[0:7]
	}
}

func NewGame(deck *Deck) *Game {
	l := MakeLibrary(deck)
	l.Shuffle()
	return &Game{
		Turn:         0,
		Life:         20,
		OpponentLife: 20,
		First:        true,
		Hand:         l[0:7],
		Library:      l[7:],
		BattleField:  nil,
	}
}

func Stats(trial int) {
	num := make(map[Status]int)
	turns := make(map[Status]int)
	for i := 0; i < trial; i++ {
		g := NewGame(MarduWorrier)
		for {
			s := g.PlayOneTurn()
			if s != Playing {
				num[s]++
				turns[s] += g.Turn
				break
			}
		}
	}
	fmt.Printf("Win: %d, Lose: %d, Draw: %d\n", num[Win], num[Lose], num[Draw])
	fmt.Printf("%f\n", float64(turns[Win])/float64(num[Win]))
}

func main() {
	g := NewGame(MarduWorrier)
	g.Print()
	for {
		s := g.PlayOneTurn()
		fmt.Println()
		g.Print()
		if s != Playing {
			break
		}
	}

	Stats(1000)
}
