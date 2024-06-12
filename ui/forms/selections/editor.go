package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	Editor struct {
		widget.BaseWidget
		characters *Characters
	}
)

func NewEditor() *Editor {
	s := &Editor{
		characters: NewCharacters(),
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Editor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s)
}
