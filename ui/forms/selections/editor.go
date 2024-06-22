package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/ff1/consts"
	"pixel-remastered-save-editor/ui/forms/editors/mapData"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Editor struct {
		widget.BaseWidget
		game      global.Game
		save      *core.Save
		allSearch *inputs.Search
	}
)

func NewEditor(game global.Game, save *core.Save) *Editor {
	s := &Editor{
		game:      game,
		save:      save,
		allSearch: inputs.NewSearch(consts.Items, consts.Weapons, consts.Shields, consts.Armors, consts.Helmets, consts.Gloves),
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Editor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewAppTabs(
			container.NewTabItem("Characters", NewCharacters(s.game, s.save, s.allSearch)),
			container.NewTabItem("Inventory", NewInventory(s.save, s.allSearch)),
			container.NewTabItem("Map Data", mapData.NewCore(s.save.Map)),
		))
}
