package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Commands struct {
		widget.BaseWidget
		enabled binding.Bool
		inputs  []*inputs.IdEntry
	}
)

func NewCoreCommands(c *core.Character) *Commands {
	e := &Commands{
		enabled: binding.BindBool(&c.EnableCommandsSave),
		inputs:  make([]*inputs.IdEntry, len(c.Commands)),
	}
	e.ExtendBaseWidget(e)
	for i := 0; i < len(c.Commands); i++ {
		j := inputs.NewIdEntryWithDataWithHint(&c.Commands[i], finder.Commands)
		j.ID.Disable()
		e.inputs[i] = j
	}
	e.enabled.AddListener(e)
	return e
}

func (e *Commands) DataChanged() {
	enabled, _ := e.enabled.Get()
	for _, i := range e.inputs {
		if enabled {
			i.ID.Enable()
		} else {
			i.ID.Disable()
		}
	}
}

func (e *Commands) CreateRenderer() fyne.WidgetRenderer {
	rows := container.NewVBox()
	for _, i := range e.inputs {
		rows.Add(container.NewGridWithColumns(3, i.Label, i.ID))
	}
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(2,
			container.NewBorder(
				widget.NewCheckWithData("Enabled", e.enabled),
				nil, nil, nil,
				container.NewVScroll(rows)))))
}
