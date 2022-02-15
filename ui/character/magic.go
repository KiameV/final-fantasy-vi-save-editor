package character

import (
	"ffvi_editor/models/consts"
	"ffvi_editor/ui"
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

	ui.InitReadOnlyText(&u.helpTB, helpText)
	return u
}

func (u *magicUI) Draw(w *nucular.Window) {
	var (
		x = 0
		y = 0
	)
	u.lastCount = 3
	w.Row(u.yLast).SpaceBegin(len(consts.Spells) + 3)

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
			u.addMagicPropertyInt(w, x*180, y, s.Name, &s.Value)
			x++
			u.lastCount++
			if x == 2 {
				x = 0
				y += 26
			}
		}
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

func (u *magicUI) Update() {

}

func (u *magicUI) addMagicPropertyInt(w *nucular.Window, x int, y int, name string, value *int) {
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
