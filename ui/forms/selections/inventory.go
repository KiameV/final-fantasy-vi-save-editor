package selections

import (
	"ffvi_editor/ui/forms/editors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Inventory struct {
		widget.BaseWidget
	}
)

func NewInventory() *Inventory {
	s := &Inventory{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Inventory) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewAppTabs(
			container.NewTabItem("Inventory", editors.NewInventory()),
			container.NewTabItem("Important", editors.NewInventoryImportant())))
}
