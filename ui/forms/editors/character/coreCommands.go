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
		job     *inputs.IdEntry
	}
)

func NewCoreCommands(c *core.Character) *Commands {
	e := &Commands{
		enabled: binding.BindBool(&c.EnableCommandsSave),
		inputs:  make([]*inputs.IdEntry, len(c.Commands)),
		job:     inputs.NewIdEntryWithDataWithHint(&c.Base.JobID, finder.Jobs),
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
	if enabled {
		e.job.ID.Enable()
	} else {
		e.job.ID.Disable()
	}
}

func (e *Commands) CreateRenderer() fyne.WidgetRenderer {
	rows := container.NewVBox()
	rows.Add(container.NewGridWithColumns(3, widget.NewLabel("Job:"), e.job.ID, e.job.Label))
	for _, i := range e.inputs {
		rows.Add(container.NewGridWithColumns(3, i.Label, i.ID))
	}
	cmdSearch := inputs.GetSearches().Commands
	jobSearch := inputs.GetSearches().Jobs
	top := container.NewVBox(
		widget.NewLabel("can cause soft locks when loading save"),
		widget.NewCheckWithData("Enabled", e.enabled))
	left := container.NewVScroll(rows)
	right := container.NewGridWithColumns(4, cmdSearch.Fields(), cmdSearch.Filter(), jobSearch.Fields(), jobSearch.Filter())
	return widget.NewSimpleRenderer(container.NewBorder(top, nil, left, right))
}
