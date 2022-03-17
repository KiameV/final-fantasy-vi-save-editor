package snes

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/snes"
)

var Characters []*models.Character

func init() {
	Characters = make([]*models.Character, len(snes.Characters))
	defaultCommand := consts.CommandLookupByValue[0xFF]
	for i, name := range snes.Characters {
		c := &models.Character{
			RootName:      name,
			Name:          name,
			StatusEffects: consts.NewStatusEffects(),
			Command1:      defaultCommand,
			Command2:      defaultCommand,
			Command3:      defaultCommand,
			Command4:      defaultCommand,
		}
		c.SpellsByIndex, c.SpellsSorted = NewSpells()
		Characters[i] = c
	}
}

func GetCharacter(name string) (c *models.Character) {
	for _, c = range Characters {
		if c.RootName == name {
			break
		}
	}
	return
}
