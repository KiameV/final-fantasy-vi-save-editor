package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
	"github.com/aarzilli/nucular"
)

type commandUI struct {
}

func newCommandUI() widget {
	return &commandUI{}
}

func (u *commandUI) Draw(w *nucular.Window) {
	var (
		commandsSorted              []string
		commandsLookupBySortedIndex []*models.Command
	)

	if global.IsShowingPR() {
		commandsSorted = pr.CommandsSorted
		commandsLookupBySortedIndex = pr.CommandsLookupBySortedIndex
	} else {
		commandsSorted = snes.CommandsSorted
		commandsLookupBySortedIndex = snes.CommandsLookupBySortedIndex
	}

	if global.IsShowingPR() {
		w.Row(24).Static(260)
		w.CheckboxText("Enable", &character.EnableCommandsSave)
		w.Row(24).Static(260)
		w.Label("Can cause soft locks and crashing!", "LC")
		w.Row(10).Static()
	}

	if character.EnableCommandsSave {
		for i, c := range character.Commands {
			w.Row(18).Static(100)
			if j := w.ComboSimple(commandsSorted, c.SortedIndex, 12); j != c.SortedIndex {
				character.Commands[i] = commandsLookupBySortedIndex[j]
			}
		}
	}
}

func (u *commandUI) Update(character *models.Character) {

}
