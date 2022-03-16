package skills

import (
	"ffvi_editor/global"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

type skillsUI struct{}

func NewUI() ui.UI {
	return &skillsUI{}
}

func (u *skillsUI) Draw(w *nucular.Window) {
	if global.IsShowingPR() {
		u.drawPR(w)
	} else {
		u.drawSnes(w)
	}
}

func (u *skillsUI) drawPR(w *nucular.Window) {
	i := 0
	/*	w.Row(18).Static(150, 10, 150, 10, 150)
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
	*/
	w.Row(18).Static(100, 20, 100)
	w.Label("Rage", "LC")
	w.Spacing(1)
	if w.ButtonText("Select All") {
		for _, r := range pr.Rages {
			r.Checked = true
		}
	}

	fourth := len(pr.SortedRages) / 4
	w.Row(1240).Static(150, 10, 150, 10, 150, 10, 150)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for i = 0; i <= fourth; i++ {
			r := pr.SortedRages[i]
			sw.CheckboxText(r.Name, &r.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*2+1; i++ {
			r := pr.SortedRages[i]
			sw.CheckboxText(r.Name, &r.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*3+2; i++ {
			r := pr.SortedRages[i]
			sw.CheckboxText(r.Name, &r.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i < len(pr.SortedRages); i++ {
			r := pr.SortedRages[i]
			sw.CheckboxText(r.Name, &r.Checked)
		}
		sw.GroupEnd()
	}
}

func (u *skillsUI) drawSnes(w *nucular.Window) {
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

	fourth := len(snes.SortedRages) / 4
	w.Row(1240).Static(150, 10, 150, 10, 150, 10, 150)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for i = 0; i <= fourth; i++ {
			se = snes.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*2+1; i++ {
			se = snes.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i <= fourth*3+2; i++ {
			se = snes.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
	w.Spacing(1)
	if sw := w.GroupBegin("", nucular.WindowBorder|nucular.WindowNoScrollbar); sw != nil {
		sw.Row(18).Static(100)
		for ; i < len(snes.SortedRages); i++ {
			se = snes.SortedRages[i]
			sw.CheckboxText(se.Name, &se.Checked)
		}
		sw.GroupEnd()
	}
}

func (u *skillsUI) Refresh() {

}

func (u *skillsUI) Name() string {
	return "Skills"
}

func (u *skillsUI) IsPRSupported() bool {
	return true
}
