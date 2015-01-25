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
	Type Type
	Name string
	Cost []Mana
}

var ChiefOfTheEdge = &Card{
	Type: Creature,
	Name: "Chief of the Edge",
	Cost: []Mana{White, Black},
}

var Swamp = &Card{
	Type: Land,
	Name: "Swamp",
}

var Plains = &Card{
	Type: Land,
	Name: "Plains",
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

type Game struct {
	Life    int
	Hand    []*Card
	Library []*Card
}

func (g *Game) Print() {
	fmt.Printf("Life: %d\n", g.Life)
	fmt.Printf("Hand (%d):\n", len(g.Hand))
	for i, c := range g.Hand {
		fmt.Printf("%d: %s\n", i, c.Name)
	}
	fmt.Printf("Library (%d):\n", len(g.Library))
	for i, c := range g.Library {
		fmt.Printf("%d: %s\n", i, c.Name)
	}
}

func NewGame(deck *Deck) *Game {
	l := MakeLibrary(deck)
	l.Shuffle()
	return &Game{
		Life:    20,
		Hand:    l[0:7],
		Library: l[7:],
	}
}

func main() {
	g := NewGame(MarduWorrier)
	g.Print()
}
