package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Equipment struct {
		widget.BaseWidget
		inputs []*inputs.IdEntry
		search *inputs.Search
	}
)

func NewCoreEquipment(c *core.Character, search *inputs.Search) *Equipment {
	e := &Equipment{
		inputs: make([]*inputs.IdEntry, len(c.Equipment)),
		search: search,
	}
	e.ExtendBaseWidget(e)
	for i, j := range c.Equipment {
		e.inputs[i] = inputs.NewIdEntryWithDataWithHint(&j.ContentID, finder.Items)
	}
	return e
}

func (e *Equipment) CreateRenderer() fyne.WidgetRenderer {
	rows := container.NewVBox()
	for _, i := range e.inputs {
		rows.Add(container.NewGridWithColumns(3, i.Label, i.ID))
	}
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(3,
			container.NewVScroll(rows),
			e.search.Fields(),
			e.search.Filter())))
}
