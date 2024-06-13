package editors

import (
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	Skills struct {
		widget.BaseWidget
		bushidos fyne.CanvasObject
		blitzes  fyne.CanvasObject
		dances   fyne.CanvasObject
		lores    fyne.CanvasObject
		rages    fyne.CanvasObject
	}
	skillSet struct {
		widget.BaseWidget
		options   []*consts.NameValueChecked
		container *fyne.Container
		checks    []binding.Bool
	}
)

func NewSkills() *Skills {
	e := &Skills{
		bushidos: newCheckGroup(pr.Bushidos),
		blitzes:  newCheckGroup(pr.Blitzes),
		dances:   newCheckGroup(pr.Dances),
		lores:    newCheckGroup(pr.Lores),
		rages:    newCheckGroup(pr.Rages),
	}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Skills) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewGridWithColumns(3,
		container.NewVScroll(container.NewVBox(
			widget.NewCard("Bushido", "", e.bushidos),
			widget.NewCard("Blitz", "", e.blitzes),
			widget.NewCard("Dance", "", e.dances),
		)),
		container.NewVScroll(container.NewVBox(
			widget.NewCard("Lore", "", e.lores))),
		container.NewVScroll(container.NewVBox(
			widget.NewCard("Rage", "", e.rages))),
	))
}

func newCheckGroup(options []*consts.NameValueChecked) *skillSet {
	s := &skillSet{
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

func (e *skillSet) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.container)
}

func (e *skillSet) selectAll() {
	for _, c := range e.checks {
		_ = c.Set(true)
	}
}

func (e *skillSet) deselectAll() {
	for _, c := range e.checks {
		_ = c.Set(false)
	}
}
