package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Veldt struct {
		widget.BaseWidget
	}
)

func NewVeldt() *Veldt {
	e := &Veldt{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Veldt) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
