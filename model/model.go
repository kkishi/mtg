package model

import (
	"github.com/kkishi/mtg/ability"
)

type Part int

const (
	// Beginning phase.
	BeginningPhase Part = iota
	UntapStep
	UpkiipStep
	DrawStep

	// Main phase.
	FirstMainPhase

	// Combat phase.
	CombatPhase
	BeginningOfCombatStep
	DeclateAtackersStep
	DeclareBlockersStep
	CombatDamageStep
	EndOfCombatStep

	// Main phase.
	SecondMainPhase

	// Ending phase.
	EndStep
	CleanupStep
)

type Game struct {
	Players      []*Player
	ActivePlayer int
	CurrentPart  Part
}

type Player struct {
	Turn        int
	First       bool
	Life        int
	Library     []*Card
	Hand        []*Card
	BattleField []*Permanent
	GraveYard   []*Card
	ManaPool    []Mana
}

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
	// Card types.
	Artifact Type = iota
	Creature
	Enchantment
	Instant
	Land
	Planeswalker
	Socery

	// Creature types.
	Human
	Worrier
	Demon
	Cat

	// Basic land types.
	Plains
	Island
	Swamp
	Mountain
	Forest
)

type Card struct {
	Name               string
	Type               Type
	SubTypes           []Type
	Cost               []Mana
	Power              int
	Toughness          int
	ActivatedAbilities []ability.ActivatedAbility
}

type Permanent struct {
	Type   Type
	Card   *Card
	Tapped bool
}
