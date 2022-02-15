package character

import (
	"ffvi_editor/models/consts"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

type statusEffectsUI struct {
	yLast int
}

func newStatusEffectsUI() widget {
	return &statusEffectsUI{}
}

func (u *statusEffectsUI) Draw(w *nucular.Window) {
	w.Row(u.yLast).SpaceBegin(len(consts.StatusEffects) + 1)
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: 0,
		W: 100,
		H: 22,
	})
	y := 24
	w.Label("Status Effects:", "LC")
	for _, se := range character.StatusEffects {
		w.LayoutSpacePush(rect.Rect{
			X: 5,
			Y: y,
			W: 80,
			H: 22,
		})
		w.CheckboxText(se.Name, &se.Checked)
		y += 24
	}
	u.yLast = y
}

func (u *statusEffectsUI) Update() {

}
