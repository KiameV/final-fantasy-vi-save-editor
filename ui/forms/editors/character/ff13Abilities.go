package character

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	FF13Abilities struct {
		widget.BaseWidget
		c         *core.Character
		abilities *fyne.Container
		body      *fyne.Container
		add       *widget.Button
	}
)

func NewFF13Abilities(c *core.Character) *FF13Abilities {
	e := &FF13Abilities{
		c:         c,
		body:      container.NewVBox(),
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)

	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})

	for i, asd := range e.c.AbilitySlotData {
		e.body.Add(widget.NewLabel(fmt.Sprintf("%d", i+1)))
		for _, a := range asd.SlotInfo.Values {
			r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
			e.body.Add(container.NewPadded(container.NewGridWithColumns(3, r.Label, r.ID)))
		}
	}
	e.populate()
	return e
}

func (e *FF13Abilities) addAbility() {
	a := &save.Ability{}
	e.c.AddAbility(a)
	e.populate()
}

func (e *FF13Abilities) removeAbility(index int) {
	e.c.RemoveAbility(index)
	e.populate()
}

func (e *FF13Abilities) populate() {
	e.abilities.RemoveAll()
	for i, a := range e.c.Abilities {
		func(i int, a *save.Ability) {
			r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
			e.abilities.Add(container.NewGridWithColumns(4, r.Label, r.ID, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				e.removeAbility(i)
			})))
		}(i, a)
	}
}

func (e *FF13Abilities) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Abilities
	left := container.NewBorder(
		container.NewGridWithColumns(4, container.NewStack(), widget.NewLabel("Ability ID"), e.add),
		nil, nil, nil,
		container.NewVScroll(e.abilities))
	right := container.NewGridWithColumns(2, search.Fields(), search.Filter())
	return widget.NewSimpleRenderer(
		container.NewBorder(nil, nil,
			left, right,
			container.NewBorder(nil, nil,
				container.NewVScroll(e.body), nil)))
}
