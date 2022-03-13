package consts

import (
	"sort"
)

type Command struct {
	NameValue
	SortedIndex int
}

var (
	Commands = []*Command{
		{NameValue: NewNameValue("Fight", 0x0)},
		{NameValue: NewNameValue("Item", 0x1)},
		{NameValue: NewNameValue("Magic", 0x2)},
		{NameValue: NewNameValue("Morph", 0x3)},
		{NameValue: NewNameValue("Revert", 0x4)},
		{NameValue: NewNameValue("Steal", 0x5)},
		{NameValue: NewNameValue("Capture", 0x6)},
		{NameValue: NewNameValue("SwdTech", 0x7)},
		{NameValue: NewNameValue("Throw", 0x8)},
		{NameValue: NewNameValue("Tools", 0x9)},
		{NameValue: NewNameValue("Blitz", 0x0A)},
		{NameValue: NewNameValue("Runic", 0x0B)},
		{NameValue: NewNameValue("Lore", 0x0C)},
		{NameValue: NewNameValue("Sketch", 0x0D)},
		{NameValue: NewNameValue("Control", 0x0E)},
		{NameValue: NewNameValue("Slot", 0x0F)},
		{NameValue: NewNameValue("Rage", 0x10)},
		{NameValue: NewNameValue("Leap", 0x11)},
		{NameValue: NewNameValue("Mimic", 0x12)},
		{NameValue: NewNameValue("Dance", 0x13)},
		{NameValue: NewNameValue("Row", 0x14)},
		{NameValue: NewNameValue("Def", 0x15)},
		{NameValue: NewNameValue("Jump", 0x16)},
		{NameValue: NewNameValue("XMagic", 0x17)},
		{NameValue: NewNameValue("GPRain", 0x18)},
		{NameValue: NewNameValue("Summon", 0x19)},
		{NameValue: NewNameValue("Health", 0x1A)},
		{NameValue: NewNameValue("Shock", 0x1B)},
		{NameValue: NewNameValue("Possess", 0x1C)},
		{NameValue: NewNameValue("Magitek", 0x1D)},
		{NameValue: NewNameValue("Blank", 0x1E)},
		{NameValue: NewNameValue("Unassigned", 0xFF)},
	}
	CommandsSorted              []string
	CommandsLookupBySortedIndex []*Command
	CommandLookupByName         map[string]*Command
	CommandLookupByValue        map[int]*Command
)

func init() {
	CommandLookupByValue = make(map[int]*Command)
	CommandLookupByName = make(map[string]*Command)
	CommandsSorted = make([]string, len(Commands))
	CommandsLookupBySortedIndex = make([]*Command, len(Commands))

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
