package inventory

import (
	"ffvi_editor/global"
	"ffvi_editor/models/snes"
	"ffvi_editor/ui"
	"fmt"
	"github.com/aarzilli/nucular"
)

type inventoryUI struct {
	ids       []*nucular.TextEditor
	counts    []*nucular.TextEditor
	countLast int
	yLast     int
}

func NewUI() ui.UI {
	inv := &inventoryUI{
		ids:    make([]*nucular.TextEditor, 255),
		counts: make([]*nucular.TextEditor, 255),
	}

	for i := 0; i < 255; i++ {
		tb := &nucular.TextEditor{}
		inv.ids[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true

		tb = &nucular.TextEditor{}
		inv.counts[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true
	}
	return inv
}

func (u *inventoryUI) Draw(w *nucular.Window) {
	if global.IsShowingPR() {
		u.drawPR(w)
	} else {
		u.drawSnes(w)
	}
}

func (u *inventoryUI) Refresh() {
	for i, row := range snes.GetInventoryRows() {
		ids := u.ids[i]
		counts := u.counts[i]
		if global.IsShowingPR() {
			ids.Maxlen = 3
			counts.Maxlen = 3
		} else {
			ids.Maxlen = 2
			counts.Maxlen = 2
		}
		ids.SelectAll()
		ids.DeleteSelection()
		ids.Text([]rune(row.ItemID))

		counts.SelectAll()
		counts.DeleteSelection()
		counts.Text([]rune(fmt.Sprintf("%d", row.Count)))
	}
}

func (u *inventoryUI) Name() string {
	return "Inventory"
}

func (u *inventoryUI) Behavior() ui.Behavior {
	return ui.Show
}
