package inventory

import (
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

type inventoryUI struct {
}

func NewInventoryUI() ui.UI {
	return &inventoryUI{}
}

func (u *inventoryUI) Draw(w *nucular.Window) {
	w.Row(18).SpaceBegin(255*4 + 5)

}
