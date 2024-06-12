package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Cheats struct {
		widget.BaseWidget
	}
)

func NewCheats() *Cheats {
	e := &Cheats{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Cheats) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
