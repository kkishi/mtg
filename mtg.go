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

type CreatureType int

const (
	Human CreatureType = iota
	Worrier
	Demon
)

type Card struct {
	Type         Type
	Name         string
	Cost         []Mana
	Produce      []Mana
	CreatureType []CreatureType
	Power        int
	Toughness    int
}

func (c *Card) IsCreatureType(ct CreatureType) bool {
	for _, cct := range c.CreatureType {
		if cct == ct {
			return true
		}
	}
	return false
}

func (c *Card) CanProduce(m Mana) bool {
	for _, p := range c.Produce {
		if m == Any || p == m {
			return true
		}
	}
	return false
}

var BloodsoakedChampion = &Card{
	Type:         Creature,
	Name:         "Bloodsoaked Champion",
	Cost:         []Mana{Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    1,
}

var ChiefOfTheEdge = &Card{
	Type:         Creature,
	Name:         "Chief of the Edge",
	Cost:         []Mana{White, Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        3,
	Toughness:    2,
}

var ChiefOfTheScale = &Card{
	Type:         Creature,
	Name:         "Chief of the Scale",
	Cost:         []Mana{White, Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    3,
}

var MarduSkullhunter = &Card{
	Type:         Creature,
	Name:         "Mardu Skullhunter",
	Cost:         []Mana{Any, Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    1,
}

var SeekerOfTheWay = &Card{
	Type:         Creature,
	Name:         "Seeker of the Way",
	Cost:         []Mana{Any, White},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    2,
}

var MarduHordechief = &Card{
	Type:         Creature,
	Name:         "Mardu Hordechief",
	Cost:         []Mana{Any, Any, White},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    3,
}

var WorrierToken = &Card{
	Type:         Creature,
	Name:         "Worrier Token",
	CreatureType: []CreatureType{Human, Worrier},
	Power:        1,
	Toughness:    1,
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
			Card:   BloodsoakedChampion,
			Amount: 4,
		},
		{
			Card:   ChiefOfTheEdge,
			Amount: 4,
		},
		{
			Card:   ChiefOfTheScale,
			Amount: 4,
		},
		{
			Card:   MarduSkullhunter,
			Amount: 4,
		},
		{
			Card:   SeekerOfTheWay,
			Amount: 4,
		},
		{
			Card:   MarduHordechief,
			Amount: 4,
		},
		{
			Card:   Swamp,
			Amount: 18,
		},
		{
			Card:   Plains,
			Amount: 18,
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
	var adjustment int
	for _, bc := range c.Game.BattleField {
		if bc.Card == ChiefOfTheEdge && bc != c && c.Card.IsCreatureType(Worrier) {
			adjustment++
		}
	}
	return c.Card.Power + adjustment
}

func (c *CardInPlay) Toughness() int {
	var adjustment int
	for _, bc := range c.Game.BattleField {
		if bc.Card == ChiefOfTheScale && bc != c && c.Card.IsCreatureType(Worrier) {
			adjustment++
		}
	}
	return c.Card.Toughness + adjustment
}

type Game struct {
	Turn         int
	Attacked     bool
	Life         int
	OpponentLife int
	First        bool
	Hand         []*Card
	Library      []*Card
	BattleField  []*CardInPlay
}

func (g *Game) Print() {
	fmt.Printf("Turn: %d\n", g.Turn)
	fmt.Printf("Attacked: %t\n", g.Attacked)
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
		fmt.Printf("%d: %s", i, c.Card.Name)
		if c.Card.Type == Creature {
			fmt.Printf(" [%d/%d]", c.Power(), c.Toughness())
		}
		if c.Tapped {
			fmt.Printf(" [T]")
		}
		if c.SummoningSickness {
			fmt.Printf(" [S]")
		}
		fmt.Printf("\n")
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
	g.Attacked = false
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

func (g *Game) PutCreature() bool {
	for i, c := range g.Hand {
		if c.Type == Creature {
			lands := make(map[*CardInPlay]bool)
			for _, m := range c.Cost {
				var ok bool
				for _, cip := range g.BattleField {
					if !cip.Tapped && !lands[cip] && cip.Card.CanProduce(m) {
						lands[cip] = true
						ok = true
						break
					}
				}
				if !ok {
					break
				}
			}
			if len(lands) != len(c.Cost) {
				continue
			}
			for l := range lands {
				l.Tapped = true
			}
			g.BattleField = append(g.BattleField, &CardInPlay{
				Tapped:            false,
				SummoningSickness: true,
				Card:              c,
				Game:              g,
			})
			g.Hand = Take(g.Hand, i)
			if g.Attacked && c == MarduHordechief {
				g.BattleField = append(g.BattleField, &CardInPlay{
					Tapped:            false,
					SummoningSickness: true,
					Card:              WorrierToken,
					Game:              g,
				})
			}
			return true
		}
	}
	return false
}

func (g *Game) FirstMain() Status {
	return Playing
}

func (g *Game) Combat() Status {
	for _, c := range g.BattleField {
		if c.Card.Type != Creature {
			continue
		}
		if c.Tapped || c.SummoningSickness {
			continue
		}
		c.Tapped = true
		g.OpponentLife -= c.Power()
		g.Attacked = true
	}
	if g.OpponentLife <= 0 {
		return Win
	}
	return Playing
}

func (g *Game) SecondMain() Status {
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

	// Put creatures.
	for g.PutCreature() {
	}
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
	winMax := 0
	winMin := 1000
	for i := 0; i < trial; i++ {
		g := NewGame(MarduWorrier)
		for {
			s := g.PlayOneTurn()
			if s != Playing {
				num[s]++
				turns[s] += g.Turn
				if s == Win {
					if winMax < g.Turn {
						winMax = g.Turn
					}
					if winMin > g.Turn {
						winMin = g.Turn
					}
				}
				break
			}
		}
	}
	fmt.Printf("Win: %d, Lose: %d, Draw: %d\n", num[Win], num[Lose], num[Draw])
	fmt.Printf("Avg: %f, Min: %d, Max: %d\n", float64(turns[Win])/float64(num[Win]), winMin, winMax)
}

func main() {
	for i := 0; i < 10; i++ {
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
	}

	Stats(1000)
}
