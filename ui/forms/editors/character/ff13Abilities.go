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
		bodyRight *fyne.Container
		add       *widget.Button
	}
)

func NewFF13Abilities(c *core.Character) *FF13Abilities {
	e := &FF13Abilities{
		c:         c,
		bodyRight: container.NewVBox(),
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)

	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})

	for i, asd := range e.c.AbilitySlotData {
		l := container.NewGridWithRows(4)
		for _, a := range asd.SlotInfo.Values {
			func(a *save.Ability) {
				r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
				l.Add(container.NewGridWithColumns(4, r.Label, r.ID))
			}(a)
		}
		e.bodyRight.Add(widget.NewCard(fmt.Sprintf("%d:", i+1), "", l))
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
	return widget.NewSimpleRenderer(container.NewGridWithColumns(7,
		container.NewBorder(
			container.NewGridWithColumns(4, container.NewStack(), widget.NewLabel("Ability ID"), container.NewStack(), e.add),
			nil, nil, nil,
			container.NewVScroll(e.abilities)),
		container.NewVScroll(e.bodyRight),
		container.NewStack(),
		search.Fields(), search.Filter()))
}
