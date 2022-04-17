package models

import (
	"ffvi_editor/models/consts"
)

type Character struct {
	ID        int
	RootName  string
	Name      string
	Level     int
	Exp       int
	HP        CurrentMax
	MP        CurrentMax
	Vigor     int
	Stamina   int
	Speed     int
	Magic     int
	IsEnabled bool
	IsNPC     bool
	//Esper   *Esper

	SpellsByIndex []*Spell
	SpellsSorted  []*Spell
	SpellsByID    map[int]*Spell

	Equipment Equipment

	Command1      *consts.Command
	Command2      *consts.Command
	Command3      *consts.Command
	Command4      *consts.Command
	StatusEffects []*consts.NameSlotMask8
}

type Spell struct {
	Name  string
	Index int
	Value int
}
