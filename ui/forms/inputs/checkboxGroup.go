package inputs

import (
	"ffvi_editor/models/consts"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	CheckboxGroup struct {
		widget.BaseWidget
		options   []*consts.NameValueChecked
		container *fyne.Container
		checks    []binding.Bool
	}
)

func NewCheckboxGroup(options []*consts.NameValueChecked) *CheckboxGroup {
	s := &CheckboxGroup{
		options:   options,
		container: container.NewVBox(),
		checks:    make([]binding.Bool, len(options)),
	}
	s.ExtendBaseWidget(s)

	s.container.Add(container.NewGridWithColumns(3,
		widget.NewButton("All", s.selectAll),
		widget.NewButton("None", s.deselectAll),
	))
	for i, j := range options {
		b := binding.BindBool(&j.Checked)
		s.checks[i] = b
		s.container.Add(widget.NewCheckWithData(j.Name, b))
	}
	return s
}

func (e *CheckboxGroup) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.container)
}

func (e *CheckboxGroup) selectAll() {
	for _, c := range e.checks {
		_ = c.Set(true)
	}
}

func (e *CheckboxGroup) deselectAll() {
	for _, c := range e.checks {
		_ = c.Set(false)
	}
}
