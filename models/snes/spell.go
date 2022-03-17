package snes

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts/snes"
)

func NewSpells() (byIndex []*models.Spell, alpha []*models.Spell) {
	byIndex = make([]*models.Spell, len(snes.Spells))
	alpha = make([]*models.Spell, len(snes.Spells))
	lookup := make(map[string]*models.Spell)
	for i, s := range snes.Spells {
		spell := &models.Spell{
			Name:  s,
			Index: i,
			Value: 0,
		}
		byIndex[i] = spell
		lookup[s] = spell
	}

	for i, s := range snes.SortedSpells {
		alpha[i] = lookup[s]
	}
	return
}
