package models

import (
	"ffvi_editor/models/consts"
)

type Character struct {
	Name    string
	Level   uint8
	Exp     uint32
	HP      CurrentMax
	MP      CurrentMax
	Vigor   uint8
	Stamina uint8
	Speed   uint8
	Magic   uint8

	SpellsByIndex []*Spell
	SpellsLookup  map[string]*Spell
	Equipment     Equipment

	CommandOne    consts.NameValue
	CommandTwo    consts.NameValue
	CommandThree  consts.NameValue
	CommandFour   consts.NameValue
	StatusEffects []uint8
}

var Characters []*Character

func init() {
	Characters = make([]*Character, len(consts.Characters))
	for i, name := range consts.Characters {
		c := &Character{
			Name: name,
		}
		c.SpellsByIndex, c.SpellsLookup = NewSpells()
		Characters[i] = c
	}
}

func GetCharacter(name string) (c *Character) {
	for _, c = range Characters {
		if c.Name == name {
			break
		}
	}
	return
}
