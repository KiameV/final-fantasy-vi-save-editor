package selections

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
	s := &Misc{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Misc) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s)
}
