package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Misc struct {
		widget.BaseWidget
	}
)

func NewMisc() *Misc {
	e := &Misc{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Misc) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
