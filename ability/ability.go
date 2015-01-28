package ability

import (
	"github.com/kkishi/mtg/model"
)

type Context struct {
	Game      *model.Game
	Player    *model.Player
	Permanent *model.Permanent
}

type ActivatedAbility interface {
	Commands(c *Context) []Command
}

type ManaAbility struct {
	Mana model.Mana
}

func (ma *ManaAbility) Commands(c *Context) []Command {
	return []Command{
		&MultiCommand{
			Commands: []Command{
				&TapCommand{
					Permanent: c.Permanent,
				},
				&AddManaCommand{
					Manas:  []model.Mana{ma.Mana},
					Player: c.Player,
				},
			},
		},
	}
}

type Command interface {
	Execute()
	Undo()
}

type MultiCommand struct {
	Commands []Command
}

func (mc *MultiCommand) Execute() {
	for _, c := range mc.Commands {
		c.Execute()
	}
}

func (mc *MultiCommand) Undo() {
	for i := range mc.Commands {
		mc.Commands[len(mc.Commands)-1-i].Undo()
	}
}

type TapCommand struct {
	Permanent *model.Permanent
}

func (tc *TapCommand) Execute() {
	tc.Permanent.Tapped = true
}

func (tc *TapCommand) Undo() {
	tc.Permanent.Tapped = false
}

type AddManaCommand struct {
	Manas  []model.Mana
	Player *model.Player
}

func (amc *AddManaCommand) Execute() {
	amc.Player.ManaPool = append(amc.Player.ManaPool, amc.Manas...)
}

func (amc *AddManaCommand) Undo() {
	amc.Player.ManaPool =
		amc.Player.ManaPool[0 : len(amc.Player.ManaPool)-len(amc.Manas)]
}
