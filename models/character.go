package models

import (
	"ffvi_editor/models/consts"
)

type Character struct {
	RootName string
	Name     string
	Level    int
	Exp      int
	HP       CurrentMax
	MP       CurrentMax
	Vigor    int
	Stamina  int
	Speed    int
	Magic    int
	//Esper   *Esper

	SpellsByIndex []*Spell
	SpellsSorted  []*Spell
	SpellsLookup  map[string]*Spell
	Equipment     Equipment

	Command1      *consts.Command
	Command2      *consts.Command
	Command3      *consts.Command
	Command4      *consts.Command
	StatusEffects []*consts.NameSlotMask8
}

var Characters []*Character

func init() {
	Characters = make([]*Character, len(consts.Characters))
	defaultCommand := consts.CommandLookupByValue[0xFF]
	for i, name := range consts.Characters {
		c := &Character{
			RootName:      name,
			Name:          name,
			StatusEffects: consts.NewStatusEffects(),
			Command1:      defaultCommand,
			Command2:      defaultCommand,
			Command3:      defaultCommand,
			Command4:      defaultCommand,
		}
		c.SpellsByIndex, c.SpellsSorted, c.SpellsLookup = NewSpells()
		Characters[i] = c
	}
}

func GetCharacter(name string) (c *Character) {
	for _, c = range Characters {
		if c.RootName == name {
			break
		}
	}
	return
}
