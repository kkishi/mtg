package card

import (
	"github.com/kkishi/mtg/ability"
	"github.com/kkishi/mtg/model"
)

var OreskosSwiftclaw = &model.Card{
	Name:      "Oreskos Swiftclaw",
	Type:      model.Creature,
	SubTypes:  []model.Type{model.Cat, model.Worrier},
	Cost:      []model.Mana{model.Any, model.White},
	Power:     3,
	Toughness: 1,
}

var Plains = &model.Card{
	Name:     "Plains",
	Type:     model.Land,
	SubTypes: []model.Type{model.Plains},
	ActivatedAbilities: []ability.ActivatedAbility{
		&ability.ManaAbility{
			Mana: White,
		},
	},
}
