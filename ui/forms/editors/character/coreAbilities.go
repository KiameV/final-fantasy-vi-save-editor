package character

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Abilities struct {
		widget.BaseWidget
		inputs []fyne.CanvasObject
	}
)

func NewCoreAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		inputs: make([]fyne.CanvasObject, len(c.Abilities)),
	}
	e.ExtendBaseWidget(e)
	for i, a := range c.Abilities {
		e.inputs[i] = container.NewGridWithColumns(3,
			widget.NewLabel(strconv.Itoa(a.ID)),
			widget.NewLabel(strconv.Itoa(a.ContentID)),
			inputs.NewIntEntryWithData(&a.SkillLevel))
	}
	return e
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(2,
			container.NewBorder(
				container.NewGridWithColumns(3,
					widget.NewLabel("Ability ID"),
					widget.NewLabel("Content ID"),
					widget.NewLabel("Level")),
				nil, nil, nil,
				container.NewVScroll(
					container.NewVBox(e.inputs...))))))
}
