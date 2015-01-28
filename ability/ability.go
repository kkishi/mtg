package ability

import (
	"github.com/kkishi/mtg/model"
)

// ManaAbility implements model.ActivatedAbility.
var _ model.ActivatedAbility = (*ManaAbility)(nil)

type ManaAbility struct {
	Mana model.Mana
}

func (ma *ManaAbility) Commands(c *model.Context) []model.Command {
	return []model.Command{
		&MultiCommand{
			Commands: []model.Command{
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

// MultiCommand implements model.Command.
var _ model.Command = (*MultiCommand)(nil)

type MultiCommand struct {
	Commands []model.Command
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

// TapCommand implements model.Command.
var _ model.Command = (*TapCommand)(nil)

type TapCommand struct {
	Permanent *model.Permanent
}

func (tc *TapCommand) Execute() {
	tc.Permanent.Tapped = true
}

func (tc *TapCommand) Undo() {
	tc.Permanent.Tapped = false
}

// AddManaCommand implements model.Command.
var _ model.Command = (*AddManaCommand)(nil)

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
