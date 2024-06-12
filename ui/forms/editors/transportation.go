package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Transportation struct {
		widget.BaseWidget
	}
)

func NewTransportation() *Transportation {
	e := &Transportation{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Transportation) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
