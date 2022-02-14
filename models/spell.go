package models

import "ffvi_editor/models/consts"

type Spell struct {
	Name  string
	Index int
	Value uint8
}

func NewSpells() (byIndex []*Spell, lookup map[string]*Spell) {
	byIndex = make([]*Spell, len(consts.Spells))
	lookup = make(map[string]*Spell)
	for i, s := range consts.Spells {
		spell := &Spell{
			Name:  s,
			Index: i,
			Value: 0,
		}
		byIndex[i] = spell
		lookup[s] = spell
	}
	return
}
