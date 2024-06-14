package editors

import (
	"ffvi_editor/models/consts/pr"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	Esper struct {
		widget.BaseWidget
		checkboxes []fyne.CanvasObject
	}
)

func NewEsper() *Esper {
	e := &Esper{checkboxes: make([]fyne.CanvasObject, len(pr.Espers))}
	e.ExtendBaseWidget(e)
	for i, esper := range pr.SortedEspers {
		e.checkboxes[i] = widget.NewCheckWithData(esper.Name, binding.BindBool(&esper.Checked))
	}
	return e
}

func (e *Esper) CreateRenderer() fyne.WidgetRenderer {
	half := len(e.checkboxes)/2 + 1
	return widget.NewSimpleRenderer(
		container.NewBorder(container.NewGridWithColumns(5,
			container.NewPadded(widget.NewButton("Select All", func() {
				for _, c := range e.checkboxes {
					c.(*widget.Check).SetChecked(true)
				}
			})),
			container.NewPadded(widget.NewButton("Deselect All", func() {
				for _, c := range e.checkboxes {
					c.(*widget.Check).SetChecked(false)
				}
			}))), nil, nil, nil,
			container.NewVScroll(
				container.NewGridWithColumns(5,
					container.NewVBox(e.checkboxes[:half]...),
					container.NewVBox(e.checkboxes[half:]...)))))
}
