package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Equipment struct {
		widget.BaseWidget
		inputs []fyne.CanvasObject
	}
)

func NewCoreEquipment(c *core.Character) *Equipment {
	e := &Equipment{
		inputs: make([]fyne.CanvasObject, len(c.Equipment)),
	}
	e.ExtendBaseWidget(e)
	for i, j := range c.Equipment {
		e.inputs[i] = container.NewGridWithColumns(3,
			inputs.NewIntEntryWithData(&j.ContentID),
			inputs.NewIntEntryWithData(&j.Count))
	}
	return e
}

func (e *Equipment) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(2,
			container.NewBorder(
				container.NewGridWithColumns(3,
					widget.NewLabel("Content ID"),
					widget.NewLabel("Count")),
				nil, nil, nil,
				container.NewVScroll(
					container.NewVBox(e.inputs...))))))
}
