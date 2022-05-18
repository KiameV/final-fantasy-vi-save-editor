package snes

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"sort"
)

var (
	Commands = []*models.Command{
		{NameValue: consts.NewNameValue("Fight", 0x0)},
		{NameValue: consts.NewNameValue("Item", 0x1)},
		{NameValue: consts.NewNameValue("Magic", 0x2)},
		{NameValue: consts.NewNameValue("Morph", 0x3)},
		{NameValue: consts.NewNameValue("Revert", 0x4)},
		{NameValue: consts.NewNameValue("Steal", 0x5)},
		{NameValue: consts.NewNameValue("Capture", 0x6)},
		{NameValue: consts.NewNameValue("SwdTech", 0x7)},
		{NameValue: consts.NewNameValue("Throw", 0x8)},
		{NameValue: consts.NewNameValue("Tools", 0x9)},
		{NameValue: consts.NewNameValue("Blitz", 0x0A)},
		{NameValue: consts.NewNameValue("Runic", 0x0B)},
		{NameValue: consts.NewNameValue("Lore", 0x0C)},
		{NameValue: consts.NewNameValue("Sketch", 0x0D)},
		{NameValue: consts.NewNameValue("Control", 0x0E)},
		{NameValue: consts.NewNameValue("Slot", 0x0F)},
		{NameValue: consts.NewNameValue("Rage", 0x10)},
		{NameValue: consts.NewNameValue("Leap", 0x11)},
		{NameValue: consts.NewNameValue("Mimic", 0x12)},
		{NameValue: consts.NewNameValue("Dance", 0x13)},
		{NameValue: consts.NewNameValue("Row", 0x14)},
		{NameValue: consts.NewNameValue("Def", 0x15)},
		{NameValue: consts.NewNameValue("Jump", 0x16)},
		{NameValue: consts.NewNameValue("XMagic", 0x17)},
		{NameValue: consts.NewNameValue("GPRain", 0x18)},
		{NameValue: consts.NewNameValue("Summon", 0x19)},
		{NameValue: consts.NewNameValue("Health", 0x1A)},
		{NameValue: consts.NewNameValue("Shock", 0x1B)},
		{NameValue: consts.NewNameValue("Possess", 0x1C)},
		{NameValue: consts.NewNameValue("Magitek", 0x1D)},
		{NameValue: consts.NewNameValue("Blank", 0x1E)},
		{NameValue: consts.NewNameValue("Unassigned", 0xFF)},
	}
	CommandsSorted              []string
	CommandsLookupBySortedIndex []*models.Command
	CommandLookupByName         map[string]*models.Command
	CommandLookupByValue        map[int]*models.Command
)

func init() {
	CommandLookupByValue = make(map[int]*models.Command)
	CommandLookupByName = make(map[string]*models.Command)
	CommandsSorted = make([]string, len(Commands))
	CommandsLookupBySortedIndex = make([]*models.Command, len(Commands))

	for i, c := range Commands {
		p := c
		CommandLookupByValue[c.Value] = p
		CommandLookupByName[c.Name] = p
		CommandsSorted[i] = c.Name
	}
	sort.Strings(CommandsSorted)

	for i, c := range CommandsSorted {
		command := CommandLookupByName[c]
		command.SortedIndex = i
		CommandsLookupBySortedIndex[i] = command
	}
}
