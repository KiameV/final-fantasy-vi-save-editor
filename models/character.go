package models

import (
	"ffvi_editor/models/consts"
)

type Character struct {
	Name    string
	Level   int
	Exp     int
	HP      CurrentMax
	MP      CurrentMax
	Vigor   int
	Stamina int
	Speed   int
	Magic   int

	SpellsByIndex []*Spell
	SpellsSorted  []*Spell
	SpellsLookup  map[string]*Spell
	Equipment     Equipment

	Command1      *consts.Command
	Command2      *consts.Command
	Command3      *consts.Command
	Command4      *consts.Command
	StatusEffects []*consts.StatusEffect
}

var Characters []*Character

func init() {
	Characters = make([]*Character, len(consts.Characters))
	for i, name := range consts.Characters {
		c := &Character{
			Name:          name,
			StatusEffects: consts.NewStatusEffects(),

			// TODO Remove
			Command1: consts.CommandsLookupBySortedIndex[1],
			Command2: consts.CommandsLookupBySortedIndex[1],
			Command3: consts.CommandsLookupBySortedIndex[1],
			Command4: consts.CommandsLookupBySortedIndex[1],
		}
		c.SpellsByIndex, c.SpellsSorted, c.SpellsLookup = NewSpells()
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
