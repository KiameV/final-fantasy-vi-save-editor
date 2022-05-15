package espers

import (
	"ffvi_editor/global"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
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
	if global.IsShowingPR() {
		u.drawPR(w)
	} else {
		u.drawSnes(w)
	}
}

func (u *espersUI) drawPR(w *nucular.Window) {
	w.Row(18).Static(100, 10, 100)
	if w.ButtonText("Select All") {
		for _, e := range pr.Espers {
			e.Checked = true
		}
	}
	w.Spacing(1)
	if w.ButtonText("Deselect All") {
		for _, e := range pr.Espers {
			e.Checked = false
		}
	}
	w.Row(u.yLast).SpaceBegin(len(snes.Espers))
	var (
		y    = 0
		half = len(pr.SortedEspers) / 2
		i    = 0
		se   *consts.NameValueChecked
	)

	for ; i <= half; i++ {
		w.LayoutSpacePush(rect.Rect{
			X: 5,
			Y: y,
			W: 80,
			H: 22,
		})
		se = pr.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}

	y = 0
	for ; i < len(pr.SortedEspers); i++ {
		w.LayoutSpacePush(rect.Rect{
			X: 95,
			Y: y,
			W: 80,
			H: 22,
		})
		se = pr.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}
	u.yLast = y + 24
}

func (u *espersUI) drawSnes(w *nucular.Window) {
	w.Row(u.yLast).SpaceBegin(len(snes.Espers))
	var (
		y    = 0
		half = len(snes.SortedEspers) / 2
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
		se = snes.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}

	y = 0
	for ; i < len(snes.SortedEspers); i++ {
		w.LayoutSpacePush(rect.Rect{
			X: 95,
			Y: y,
			W: 80,
			H: 22,
		})
		se = snes.SortedEspers[i]
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}
	u.yLast = y + 24
}

func (u *espersUI) Refresh() {

}

func (u *espersUI) Name() string {
	return "Espers"
}

func (u *espersUI) Behavior() ui.Behavior {
	return ui.Show
}
