package espers

import (
	"ffvi_editor/models/consts"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

type espersUI struct {
	yLast int
}

func NewUI() ui.UI {
	return &espersUI{}
}

func (u *espersUI) Draw(w *nucular.Window) {
	w.Row(u.yLast).SpaceBegin(len(consts.Espers))
	var (
		y    = 0
		half = len(consts.SortedEspers) / 2
		i    = 0
		se   *consts.NameSlotMask8
	)

	for ; i <= half; i++ {
		w.LayoutSpacePush(rect.Rect{
			X: 5,
			Y: y,
			W: 80,
			H: 22,
		})
		se = consts.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}

	y = 0
	for ; i < len(consts.SortedEspers); i++ {
		w.LayoutSpacePush(rect.Rect{
			X: 95,
			Y: y,
			W: 80,
			H: 22,
		})
		se = consts.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}
	u.yLast = y + 24
}

func (u *espersUI) Refresh() {
	
}
