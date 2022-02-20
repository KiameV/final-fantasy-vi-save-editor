package skills

import (
	"ffvi_editor/models/consts"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

type skillsUI struct{}

func NewUI() ui.UI {
	return &skillsUI{}
}

func (u *skillsUI) Draw(w *nucular.Window) {
	w.Row(18).Static(150, 10, 150, 10, 150)
	w.Label("SwdTech", "LC")
	w.Spacing(1)
	w.Label("Blitz", "LC")
	w.Spacing(1)
	w.Label("Dance", "LC")

	w.Row(170).Static(150, 10, 150, 10, 150)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for _, se := range consts.SwordTech {
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for _, se := range consts.Blitzes {
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for _, se := range consts.Dances {
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}

	w.Row(18).Static(150)
	w.Label("Lore", "LC")

	w.Row(240).Static(150, 10, 150)
	var se *consts.NameSlotMask8
	i := 0
	half := len(consts.SortedLores) / 2
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for i = 0; i < half; i++ {
			se = consts.SortedLores[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i < len(consts.SortedLores); i++ {
			se = consts.SortedLores[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}

	w.Row(18).Static(100)
	w.Label("Rage", "LC")

	fourth := len(consts.SortedRages) / 4
	w.Row(1240).Static(150, 10, 150, 10, 150, 10, 150)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for i = 0; i <= fourth; i++ {
			se = consts.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*2+1; i++ {
			se = consts.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*3+2; i++ {
			se = consts.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i < len(consts.SortedRages); i++ {
			se = consts.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
}

func (u *skillsUI) Refresh() {

}
