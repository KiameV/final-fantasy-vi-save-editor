package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	SkillType struct {
		widget.BaseWidget
	}
)

func NewSkillType() *SkillType {
	e := &SkillType{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *SkillType) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e)
}
