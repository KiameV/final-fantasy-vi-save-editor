package pr

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"sort"
)

var (
	Commands = []*models.Command{
		{NameValue: consts.NewNameValue("[none]", 4)},
		{NameValue: consts.NewNameValue("Attack", 1)},
		{NameValue: consts.NewNameValue("Defend", 2)},
		{NameValue: consts.NewNameValue("Items", 3)},
		{NameValue: consts.NewNameValue("Row", 5)},
		{NameValue: consts.NewNameValue("Skip", 6)},
		{NameValue: consts.NewNameValue("Magitek", 7)},
		{NameValue: consts.NewNameValue("Trance", 8)},
		{NameValue: consts.NewNameValue("Revert", 9)},
		{NameValue: consts.NewNameValue("Steal", 10)},
		{NameValue: consts.NewNameValue("Mug", 11)},
		{NameValue: consts.NewNameValue("Bushido", 12)},
		{NameValue: consts.NewNameValue("Throw", 13)},
		{NameValue: consts.NewNameValue("Tools", 14)},
		{NameValue: consts.NewNameValue("Blitz", 15)},
		{NameValue: consts.NewNameValue("Runic", 16)},
		{NameValue: consts.NewNameValue("Lore", 17)},
		{NameValue: consts.NewNameValue("Sketch", 18)},
		{NameValue: consts.NewNameValue("Control", 19)},
		{NameValue: consts.NewNameValue("Slot", 20)},
		{NameValue: consts.NewNameValue("Gil Toss", 21)},
		{NameValue: consts.NewNameValue("Dance", 22)},
		{NameValue: consts.NewNameValue("Rage", 23)},
		{NameValue: consts.NewNameValue("Leap", 24)},
		{NameValue: consts.NewNameValue("Mimic", 25)},
		{NameValue: consts.NewNameValue("Magic", 26)},
		{NameValue: consts.NewNameValue("Pray", 27)},
		{NameValue: consts.NewNameValue("Shock", 28)},
		{NameValue: consts.NewNameValue("Possess", 29)},
		{NameValue: consts.NewNameValue("Jump", 30)},
		{NameValue: consts.NewNameValue("Dualcast", 31)},
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
	sort.Strings(CommandsSorted[1:])

	for i, c := range CommandsSorted {
		command := CommandLookupByName[c]
		command.SortedIndex = i
		CommandsLookupBySortedIndex[i] = command
	}
}
