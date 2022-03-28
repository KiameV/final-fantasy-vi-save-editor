package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/snes"
	"ffvi_editor/ui/file"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

var (
	minValue   = [2]int{1, 0}
	vigorLabel = [2]string{"Vigor", "Strength"}
	speedLabel = [2]string{"Speed", "Agility"}
)

type statsUI struct {
	name       nucular.TextEditor
	levelExpTB nucular.TextEditor
	yLast      int
}

func newStatsUI() widget {
	u := &statsUI{}

	u.name.Flags = nucular.EditField
	u.name.Maxlen = 8
	u.name.SingleLine = true
	u.name.Text([]rune(snes.Characters[0]))

	widgets.InitReadOnlyText(&u.levelExpTB, consts.LevelText)
	return u
}

func (u *statsUI) Draw(w *nucular.Window) {
	var y = 0
	count := 24
	if !global.IsShowingPR() || (file.PrIO != nil && !file.PrIO.HasUnicodeNames()) {
		count = 23
	}

	w.Row(u.yLast).SpaceBegin(count)
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: 0,
		W: 50,
		H: 22,
	})
	if !global.IsShowingPR() || (file.PrIO != nil && !file.PrIO.HasUnicodeNames()) {
		w.Label("Name:", "LC")
		w.LayoutSpacePush(rect.Rect{
			X: 60,
			Y: 0,
			W: 100,
			H: 22,
		})
		u.name.Edit(w)
		t := string(u.name.Buffer)
		if t != character.Name {
			character.Name = t
		}
	}

	w.LayoutSpacePush(rect.Rect{
		X: 180,
		Y: 0,
		W: 160,
		H: 22,
	})
	if w.ButtonText("Reset Exp") {
		character.Exp = int(consts.LevelToExp[character.Level])
	}
	y += 24

	addPropertyInt(w, 0, y, "Level", &character.Level, 1, 99, 1)
	addPropertyInt(w, 180, y, "Exp", &character.Exp, 1, 2637112, 1000)
	y += 26

	addPropertyInt(w, 0, y, "Current HP", &character.HP.Current, 0, 9999, 100)
	addPropertyInt(w, 180, y, "Max HP", &character.HP.Max, 1, 9999, 100)
	y += 26

	addPropertyInt(w, 0, y, "Current MP", &character.MP.Current, 0, 999, 10)
	addPropertyInt(w, 180, y, "Max MP", &character.MP.Max, 1, 999, 10)
	y += 26

	i := 0
	if global.IsShowing(global.ShowPR) {
		i = 1
	}

	addPropertyInt(w, 0, y, vigorLabel[i], &character.Vigor, minValue[i], 255, 1)
	addPropertyInt(w, 180, y, "Stamina", &character.Stamina, minValue[i], 255, 1)
	y += 26

	addPropertyInt(w, 0, y, speedLabel[i], &character.Speed, minValue[i], 255, 1)
	addPropertyInt(w, 180, y, "Magic", &character.Magic, minValue[i], 255, 1)
	y += 26

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 0,
		W: 200,
		H: u.yLast,
	})
	u.levelExpTB.Edit(w)

	character.HP.Fix()
	character.MP.Fix()

	u.yLast = y
}

func (u *statsUI) Update(character *models.Character) {
	u.name.SelectAll()
	u.name.DeleteSelection()
	u.name.Text([]rune(character.Name))
}

func addPropertyInt(w *nucular.Window, x int, y int, name string, value *int, min int, max int, step int) {
	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 160,
		H: 24,
	})
	w.PropertyInt(name, min, value, max, step, 0)
}
