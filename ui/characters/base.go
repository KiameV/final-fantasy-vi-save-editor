package characters

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	f "ffvi_editor/ui/factory"
	"github.com/therecipe/qt/widgets"
)

var (
	selectedCharacter *models.Character
	statsLayout       *widgets.QWidget
)

func CreateCharactersLayout(window *widgets.QMainWindow) widgets.QWidget_ITF {
	g := widgets.NewQVBoxLayout()

	selectedCharacter = models.GetCharacter(consts.Characters[0])
	sl := createStatsLayout()
	ml := createMagicLayout()
	updateStats()

	g.AddWidget(f.CreateComboInput("Character", consts.Characters, func(name string) {
		selectedCharacter = models.GetCharacter(name)
		updateStats()
		updateMagic()
	}), 0, 0)

	t := widgets.NewQTabWidget(nil)
	t.AddTab(sl, "Stats")
	t.AddTab(ml, "Magic")
	//t.AddTab(el, "Equipment")
	//t.AddTab(cl, "Commands")
	//t.AddTab(sel, "Status Effects")

	g.AddWidget(t, 0, 0)

	w := widgets.NewQWidget(window, 0)
	w.SetLayout(g)
	return w
}
