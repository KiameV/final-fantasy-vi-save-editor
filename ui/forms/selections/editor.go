package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/editors/mapData"
)

type (
	Editor struct {
		widget.BaseWidget
		save *core.Save
	}
)

func NewEditor(save *core.Save) *Editor {
	s := &Editor{save: save}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Editor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewAppTabs(
			container.NewTabItem("Characters", NewCharacters(s.save)),
			container.NewTabItem("Inventory", NewInventory(s.save)),
			container.NewTabItem("Map Data", mapData.NewCore(s.save.Map)),
		))
}
