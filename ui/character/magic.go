package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"strings"
)

type magicUI struct {
	searchTB  nucular.TextEditor
	helpTB    nucular.TextEditor
	yLast     int
	lastCount int
}

func newMagicUI() widget {
	u := &magicUI{}

	u.searchTB.Flags = nucular.EditField
	u.searchTB.Maxlen = 10
	u.searchTB.SingleLine = true

	widgets.InitReadOnlyText(&u.helpTB, helpText)
	return u
}

func (u *magicUI) Draw(w *nucular.Window) {
	if global.IsShowingPR() {
		u.drawPR(w)
	} else {
		u.drawSnes(w)
	}
}

func (u *magicUI) drawPR(w *nucular.Window) {
	var (
		x = 0
		y = 0
	)
	u.lastCount = 5
	w.Row(u.yLast).SpaceBegin(len(pr.Spells) + 3)

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: 0,
		W: 50,
		H: 22,
	})
	w.Label("Search:", "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 60,
		Y: 0,
		W: 100,
		H: 22,
	})
	u.searchTB.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 180,
		Y: 0,
		W: 100,
		H: 22,
	})
	if w.ButtonText("Learn All") {
		for _, s := range character.SpellsSorted {
			s.Value = 100
		}
	}
	w.LayoutSpacePush(rect.Rect{
		X: 290,
		Y: 0,
		W: 100,
		H: 22,
	})
	if w.ButtonText("Unlearn All") {
		for _, s := range character.SpellsSorted {
			s.Value = 0
		}
	}
	y += 26

	search := strings.ToLower(strings.TrimSpace(string(u.searchTB.Buffer)))
	for _, s := range character.SpellsSorted {
		if search == "" || strings.Contains(strings.ToLower(s.Name), search) {
			u.addPrMagicPropertyInt(w, x*180, y, s.Name, &s.Value)
			x++
			u.lastCount++
			if x == 2 {
				x = 0
				y += 26
			}
		}
	}
	if search != "" {
		y += 26
	}

	u.yLast = y
}

func (u *magicUI) drawSnes(w *nucular.Window) {
	var (
		x = 0
		y = 0
	)
	u.lastCount = 3
	w.Row(u.yLast).SpaceBegin(len(snes.Spells) + 3)

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: 0,
		W: 50,
		H: 22,
	})
	w.Label("Search:", "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 60,
		Y: 0,
		W: 100,
		H: 22,
	})
	u.searchTB.Edit(w)
	y += 24

	search := strings.ToLower(strings.TrimSpace(string(u.searchTB.Buffer)))
	for _, s := range character.SpellsSorted {
		if search == "" || strings.Contains(strings.ToLower(s.Name), search) {
			u.addSnesMagicPropertyInt(w, x*180, y, s.Name, &s.Value)
			x++
			u.lastCount++
			if x == 2 {
				x = 0
				y += 26
			}
		}
	}
	if search != "" {
		y += 26
	}

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 0,
		W: 200,
		H: y,
	})
	u.helpTB.Edit(w)

	u.yLast = y
}

func (u *magicUI) Update(character *models.Character) {
	u.searchTB.Text([]rune(""))
}

func (u *magicUI) addPrMagicPropertyInt(w *nucular.Window, x int, y int, name string, value *int) {
	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 160,
		H: 24,
	})
	w.PropertyInt(name, 0, value, 100, 1, 0)
}

func (u *magicUI) addSnesMagicPropertyInt(w *nucular.Window, x int, y int, name string, value *int) {
	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 160,
		H: 24,
	})
	o := *value
	w.PropertyInt(name, 0, value, 255, 1, 0)
	n := *value
	if o != n {
		if o == 255 {
			if n > 99 {
				*value = 99
			}
		} else if o >= 0 || o <= 99 {
			if n > 99 && n < 255 {
				*value = 255
			}
		}
	}
}

const helpText = `Values:
- 0-99: percent learned
- 255: spell is known`
