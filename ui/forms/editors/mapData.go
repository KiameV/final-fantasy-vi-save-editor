package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	MapData struct {
		widget.BaseWidget
	}
)

func NewMapData() *MapData {
	e := &MapData{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *MapData) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
