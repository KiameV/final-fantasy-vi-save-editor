package selections

import (
	"ffvi_editor/ui/forms/editors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Editor struct {
		widget.BaseWidget
	}
)

func NewEditor() *Editor {
	s := &Editor{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Editor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewAppTabs(
			container.NewTabItem("Characters", NewCharacters()),
			container.NewTabItem("Inventory", NewInventory()),
			container.NewTabItem("Skills", editors.NewSkills()),
		))
}
