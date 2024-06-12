package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	InventoryType struct {
		widget.BaseWidget
	}
)

func NewInventoryType() *InventoryType {
	e := &InventoryType{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *InventoryType) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
