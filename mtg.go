package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/kkishi/mtg/card"
	"github.com/kkishi/mtg/model"
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
	Cat
	Orc
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

var TormentedHero = &Card{
	Type:         Creature,
	Name:         "Tormented Hero",
	Cost:         []Mana{Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    1,
}

var MarduShadowspear = &Card{
	Type:         Creature,
	Name:         "Mardu Shadowspear",
	Cost:         []Mana{Black},
	CreatureType: []CreatureType{Human, Worrier},
	Power:        2,
	Toughness:    1,
}

var MarduWoeReaper = &Card{
	Type:         Creature,
	Name:         "Mardu Woe-Reaper",
	Cost:         []Mana{White},
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

var OreskosSwiftclaw = &Card{
	Type:         Creature,
	Name:         "OreskosSwiftclaw",
	Cost:         []Mana{Any, White},
	CreatureType: []CreatureType{Cat, Worrier},
	Power:        3,
	Toughness:    1,
}

var BattleBrawler = &Card{
	Type:         Creature,
	Name:         "Battle Brawler",
	Cost:         []Mana{Any, Black},
	CreatureType: []CreatureType{Orc, Worrier},
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

var ButcherOfTheHorde = &Card{
	Type:         Creature,
	Name:         "Butcher of the Horde",
	Cost:         []Mana{Any, Red, White, Black},
	CreatureType: []CreatureType{Demon},
	Power:        5,
	Toughness:    4,
}

var MarduCharm = &Card{
	Type:      Instant,
	Name:      "Mardu Charm",
	Cost:      []Mana{Red, White, Black},
	Power:     0,
	Toughness: 0,
}

var RaidersSpoils = &Card{
	Type:      Enchantment,
	Name:      "Raider's Spoils",
	Cost:      []Mana{Any, Any, Any, Black},
	Power:     0,
	Toughness: 0,
}

var WorrierToken = &Card{
	Type:         Creature,
	Name:         "Worrier Token",
	CreatureType: []CreatureType{Human, Worrier},
	Power:        1,
	Toughness:    1,
}

var NomadOutpost = &Card{
	Type:    Land,
	Name:    "Nomad Outpost",
	Produce: []Mana{Red, White, Black},
}

var ScouredBarrens = &Card{
	Type:    Land,
	Name:    "Scoured Barrens",
	Produce: []Mana{Black, White},
}

var CavesOfKoilos = &Card{
	Type:    Land,
	Name:    "Caves of Koilos",
	Produce: []Mana{Black, White},
}

var WindScarredCrag = &Card{
	Type:    Land,
	Name:    "Wind-Scarred Crag",
	Produce: []Mana{Red, White},
}

var BattlefieldForge = &Card{
	Type:    Land,
	Name:    "Battlefield Forge",
	Produce: []Mana{Red, White},
}

var BloodfellCaves = &Card{
	Type:    Land,
	Name:    "Bloodfell Caves",
	Produce: []Mana{Black, Red},
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
		{BloodsoakedChampion, 4},
		{TormentedHero, 4},
		{MarduWoeReaper, 4},
		{BattleBrawler, 4},
		{ChiefOfTheEdge, 4},
		{ChiefOfTheScale, 4},
		{MarduHordechief, 4},
		{MarduCharm, 4},
		{RaidersSpoils, 4},
		{CavesOfKoilos, 4},
		{Plains, 8},
		{Swamp, 12},
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
		if bc.Card == RaidersSpoils {
			adjustment++
		}
	}
	if c.Card == BattleBrawler {
	out:
		for _, bc := range c.Game.BattleField {
			for _, m := range bc.Card.Cost {
				if m == White || m == Red {
					adjustment++
					break out
				}
			}
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

	// Convenient data for search.
	BestTurn int
	BestHand []*Card
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
		if i == 10 {
			fmt.Println("...")
			break
		}
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

func (g *Game) PlayOneTurn(greedy bool) Status {
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
	if greedy {
		if s := g.MainGreedy(); s != Playing {
			return s
		}
	} else {
		if s := g.SecondMain(); s != Playing {
			return s
		}
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

type Key struct {
	Any   int
	White int
	Blue  int
	Black int
	Red   int
	Green int
	Used  int64
}

func (k *Key) Add(m Mana) {
	switch m {
	case Any:
		k.Any++
	case White:
		k.White++
	case Blue:
		k.Blue++
	case Black:
		k.Black++
	case Red:
		k.Red++
	case Green:
		k.Green++
	}
}

func Check(avail, cost int, any *int) bool {
	if avail < cost {
		return false
	}
	*any += avail - cost
	return true
}

func (k *Key) Payable(c *Key) bool {
	var any int
	if !Check(k.White, c.White, &any) {
		return false
	}
	if !Check(k.Blue, c.Blue, &any) {
		return false
	}
	if !Check(k.Black, c.Black, &any) {
		return false
	}
	if !Check(k.Red, c.Red, &any) {
		return false
	}
	if !Check(k.Green, c.Green, &any) {
		return false
	}
	return any >= c.Any
}

func (k *Key) Total() int {
	return k.Any + k.White + k.Blue + k.Black + k.Red + k.Green
}

func (g *Game) CastSpells() {
	dp := make(map[Key]bool)
	dp[Key{}] = true
	for i, cip := range g.BattleField {
		if cip.Tapped || cip.Card.Type != Land {
			continue
		}
		ndp := make(map[Key]bool)
		for key := range dp {
			ndp[key] = true
			for _, mana := range cip.Card.Produce {
				nkey := key
				nkey.Add(mana)
				nkey.Used |= 1 << uint(i)
				ndp[nkey] = true
			}
		}
		dp = ndp
	}
	for i := len(g.Hand) - 1; i >= 0; i-- {
		var cost Key
		for j := 0; j <= i; j++ {
			c := g.Hand[j]
			if c.Type != Creature && c.Type != Enchantment {
				continue
			}
			for _, m := range c.Cost {
				cost.Add(m)
			}
		}
		minPay := math.MaxInt32
		var minPayKey Key
		for k := range dp {
			if k.Payable(&cost) && k.Total() < minPay {
				minPay = k.Total()
				minPayKey = k
			}
		}
		if minPay < math.MaxInt32 {
			for i, cip := range g.BattleField {
				if (minPayKey.Used & (1 << uint(i))) != 0 {
					cip.Tapped = true
				}
			}
			var newHand []*Card
			for j, c := range g.Hand {
				if j <= i && (c.Type == Creature || c.Type == Enchantment) {
					if c.Type == Creature {
						var Tapped bool
						if c == TormentedHero || c == MarduSkullhunter {
							Tapped = true
						}
						g.BattleField = append(g.BattleField, &CardInPlay{
							Tapped:            Tapped,
							SummoningSickness: true,
							Card:              c,
							Game:              g,
						})
						if g.Attacked && c == MarduHordechief {
							g.BattleField = append(g.BattleField, &CardInPlay{
								Tapped:            false,
								SummoningSickness: true,
								Card:              WorrierToken,
								Game:              g,
							})
						}
					} else if c.Type == Enchantment {
						g.BattleField = append(g.BattleField, &CardInPlay{
							Card: c,
							Game: g,
						})
					}
				} else {
					newHand = append(newHand, c)
				}
			}
			g.Hand = newHand
			break
		}
	}
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

func CopyCards(cs []*Card) []*Card {
	var ret []*Card
	for _, c := range cs {
		ret = append(ret, c)
	}
	return ret
}

func CopyCardInPlay(cips []*CardInPlay) []*CardInPlay {
	var ret []*CardInPlay
	for _, cip := range cips {
		var copied CardInPlay = *cip
		ret = append(ret, &copied)
	}
	return ret
}

func (g *Game) Rec(depth int, used map[int]bool, perm []*Card, hand []*Card) {
	if depth == len(hand) {
		cg := &Game{
			Turn:         g.Turn,
			Attacked:     g.Attacked,
			Life:         g.Life,
			OpponentLife: g.Life,
			First:        g.First,
			Hand:         CopyCards(perm),
			Library:      CopyCards(g.Library),
			BattleField:  CopyCardInPlay(g.BattleField),
		}

		// g was about to start the second main phase. First we finish that turn.
		cg.MainGreedy()
		cg.Discard()

		for i := 0; i < g.BestTurn; i++ {
			if cg.PlayOneTurn(true) != Playing {
				g.BestTurn = i
				g.BestHand = CopyCards(perm)
			}
		}
	} else {
		for i, c := range hand {
			if used[i] {
				continue
			}
			used[i] = true
			perm = append(perm, c)
			g.Rec(depth+1, used, perm, hand)
			perm = perm[0 : len(perm)-1]
			used[i] = false
		}
	}
}

func (g *Game) MainGreedy() Status {
	// Put a first land in hand.
	for i, c := range g.Hand {
		if c.Type != Land {
			continue
		}

		var Tapped bool
		if c == NomadOutpost || c == ScouredBarrens || c == WindScarredCrag ||
			c == BloodfellCaves {
			Tapped = true
		}

		g.BattleField = append(g.BattleField, &CardInPlay{
			Tapped:            Tapped,
			SummoningSickness: false,
			Card:              c,
			Game:              g,
		})
		g.Hand = Take(g.Hand, i)
		break
	}

	// Cast as much spells as possible.
	g.CastSpells()

	return Playing
}

func (g *Game) SecondMain() Status {
	g.BestTurn = math.MaxInt32
	g.BestHand = nil
	g.Rec(0, make(map[int]bool), nil, g.Hand)
	// fmt.Printf("Best (%d): ", bestTurn)
	// for i, c := range bestHand {
	// 	if i > 0 {
	// 		fmt.Printf(", ")
	// 	}
	// 	fmt.Printf("%s", c.Name)
	// }
	// fmt.Printf("\n")
	g.Hand = g.BestHand
	return g.MainGreedy()
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

type Ints []int

func (is Ints) Len() int           { return len(is) }
func (is Ints) Less(i, j int) bool { return is[i] < is[j] }
func (is Ints) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

func Stats(trial, parallelism int) {
	turn := make(chan int)
	for i := 0; i < parallelism; i++ {
		go func() {
			for i := 0; i < trial/parallelism; i++ {
				g := NewGame(MarduWorrier)
				for {
					if g.PlayOneTurn(false) != Playing {
						turn <- g.Turn
						break
					}
				}
			}
		}()
	}

	var turns int
	var is Ints
	for i := 0; i < trial; i++ {
		t := <-turn
		turns += t
		is = append(is, t)
		fmt.Printf("Trial %d: %d turns\n", i, t)
	}
	sort.Sort(is)

	fmt.Printf("Avg: %f, 50%%: %d, 75%%: %d, 90%%: %d, 95%%: %d\n",
		float64(turns)/float64(trial), is[len(is)/2], is[len(is)*3/4],
		is[len(is)*9/10], is[len(is)*19/20])

	t := -1
	for i, it := range is {
		if t < it {
			t = it
			fmt.Printf("T%d: %.1f%%\n", t-1, float64(i)*100/float64(trial))
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		g := NewGame(MarduWorrier)
		g.Print()
		fmt.Println()
		for {
			s := g.PlayOneTurn(false)
			g.Print()
			fmt.Println()
			if s != Playing {
				break
			}
		}
	}

	Stats(100, 4)

	g := &model.Game{
		Players: []*model.Player{
			&model.Player{},
		},
	}

	p := &model.Permanent{
		Type:   model.Land,
		Card:   card.Plains,
		Tapped: false,
	}

	for _, aa := range p.Card.ActivatedAbilities {
		cs := aa.Commands(&model.Context{
			Game:      g,
			Player:    g.Players[0],
			Permanent: p,
		})
		for _, c := range cs {
			fmt.Println(g.Players[0], p)
			c.Execute()
			fmt.Println(g.Players[0], p)
			c.Undo()
			fmt.Println(g.Players[0], p)
		}
	}
}
