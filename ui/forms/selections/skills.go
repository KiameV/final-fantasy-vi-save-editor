package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Skills struct {
		widget.BaseWidget
	}
)

func NewSkills() *Skills {
	s := &Skills{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Skills) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s)
}
