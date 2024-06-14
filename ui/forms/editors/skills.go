package editors

import (
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
)

func NewSkills() *Skills {
	e := &Skills{
		bushidos: inputs.NewCheckboxGroup(pr.Bushidos),
		blitzes:  inputs.NewCheckboxGroup(pr.Blitzes),
		dances:   inputs.NewCheckboxGroup(pr.Dances),
		lores:    inputs.NewCheckboxGroup(pr.Lores),
		rages:    inputs.NewCheckboxGroup(pr.Rages),
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
