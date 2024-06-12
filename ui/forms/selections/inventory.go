package selections

import (
	"fyne.io/fyne/v2"
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
	return widget.NewSimpleRenderer(s)
}
