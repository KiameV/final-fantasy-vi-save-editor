package character

import (
	"ffvi_editor/models/consts"
	"github.com/aarzilli/nucular"
)

type commandUI struct {
}

func newCommandUI() widget {
	u := &commandUI{}

	return u
}

func (u *commandUI) Draw(w *nucular.Window) {
	var i int

	w.Row(18).Static(100)
	if i = w.ComboSimple(consts.CommandsSorted, character.Command1.SortedIndex, 12); i != character.Command1.SortedIndex {
		character.Command1 = consts.CommandsLookupBySortedIndex[i]
	}
	w.Row(18).Static(100)
	if i = w.ComboSimple(consts.CommandsSorted, character.Command2.SortedIndex, 12); i != character.Command2.SortedIndex {
		character.Command2 = consts.CommandsLookupBySortedIndex[i]
	}
	w.Row(18).Static(100)
	if i = w.ComboSimple(consts.CommandsSorted, character.Command3.SortedIndex, 12); i != character.Command3.SortedIndex {
		character.Command3 = consts.CommandsLookupBySortedIndex[i]
	}
	w.Row(18).Static(100)
	if i = w.ComboSimple(consts.CommandsSorted, character.Command4.SortedIndex, 12); i != character.Command4.SortedIndex {
		character.Command4 = consts.CommandsLookupBySortedIndex[i]
	}
}

func (u *commandUI) Update() {

}
