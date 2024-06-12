package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Party struct {
		widget.BaseWidget
	}
)

func NewParty() *Party {
	e := &Party{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Party) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
