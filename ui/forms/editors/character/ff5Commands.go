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
	FF5Commands struct {
		widget.BaseWidget
		enabled binding.Bool
		inputs  []*inputs.IdEntry
	}
)

func NewFF5Commands(c *core.Character) *FF5Commands {
	e := &FF5Commands{
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

func (e *FF5Commands) DataChanged() {
	enabled, _ := e.enabled.Get()
	for _, i := range e.inputs {
		if enabled {
			i.ID.Enable()
		} else {
			i.ID.Disable()
		}
	}
}

func (e *FF5Commands) CreateRenderer() fyne.WidgetRenderer {
	rows := container.NewVBox()
	for _, i := range e.inputs {
		rows.Add(container.NewGridWithColumns(3, i.Label, i.ID))
	}
	cmdSearch := inputs.GetSearches().Commands
	top := container.NewVBox(
		widget.NewLabel("can cause soft locks when loading save"),
		widget.NewCheckWithData("Enabled", e.enabled))
	left := container.NewVScroll(rows)
	right := container.NewGridWithColumns(2, cmdSearch.Fields(), cmdSearch.Filter())
	return widget.NewSimpleRenderer(container.NewBorder(top, nil, left, right))
}
