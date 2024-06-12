package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Esper struct {
		widget.BaseWidget
	}
)

func NewEsper() *Esper {
	e := &Esper{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Esper) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
