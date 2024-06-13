package editors

import (
	"slices"

	"ffvi_editor/models"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Commands struct {
		widget.BaseWidget
		selections []cmdSelect
	}
	cmdSelect struct {
		selection *widget.Select
		entry     *inputs.IntEntry
	}
)

func NewCommands(c *models.Character) *Commands {
	e := &Commands{
		selections: make([]cmdSelect, len(c.Commands)),
	}
	e.ExtendBaseWidget(e)
	options := make([]string, 0, len(cmdStoI))
	for k := range cmdStoI {
		options = append(options, k)
	}
	slices.Sort(options)
	for i, cmd := range c.Commands {
		e.selections[i] = newCmdSelect(options, cmd)
	}
	return e
}

func (e *Commands) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(
		widget.NewLabel("Warning: can cause soft locks and crashing in game."))
	for i, s := range e.selections {
		c.Add(container.NewGridWithColumns(2, s.selection, s.entry))
		if i == 3 {
			c.Add(widget.NewSeparator())
		}
	}
	return widget.NewSimpleRenderer(container.NewGridWithColumns(2, c))
}

func newCmdSelect(options []string, cmd *models.Command) (s cmdSelect) {
	s = cmdSelect{
		selection: widget.NewSelect(options, func(v string) {
			if i, found := cmdStoI[v]; found {
				cmd.Value = i
				s.entry.SetInt(i)
			}
		}),
		entry: inputs.NewIntEntry(),
	}
	s.entry.OnChanged = func(v string) {
		cmd.Value = s.entry.Int()
		if i, found := cmdItoS[cmd.Value]; found {
			s.selection.Selected = i
		} else {
			s.selection.Selected = unknownCmd
		}
		s.selection.Refresh()
	}
	s.entry.SetInt(cmd.Value)
	return s
}

func init() {
	for k, v := range cmdStoI {
		cmdItoS[v] = k
	}
}

const (
	unknownCmd = "[unknown]"
)

var (
	cmdItoS = make(map[int]string)
	cmdStoI = map[string]int{
		"[none]":   4,
		"Attack":   1,
		"Defend":   2,
		"Items":    3,
		"Row":      5,
		"Skip":     6,
		"Magitek":  7,
		"Trance":   8,
		"Revert":   9,
		"Steal":    10,
		"Mug":      11,
		"Bushido":  12,
		"Throw":    13,
		"Tools":    14,
		"Blitz":    15,
		"Runic":    16,
		"Lore":     17,
		"Sketch":   18,
		"Control":  19,
		"Slot":     20,
		"Gil Toss": 21,
		"Dance":    22,
		"Rage":     23,
		"Leap":     24,
		"Mimic":    25,
		"Magic":    26,
		"Pray":     27,
		"Shock":    28,
		"Possess":  29,
		"Jump":     30,
		"Dualcast": 31,
		unknownCmd: 0,
	}
)
