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
