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
	FF1Abilities struct {
		widget.BaseWidget
		c         *core.Character
		abilities *fyne.Container
		bodyLeft  *fyne.Container
		bodyRight *fyne.Container
		add       *widget.Button
	}
)

func NewFF1Abilities(c *core.Character) *FF1Abilities {
	e := &FF1Abilities{
		c:         c,
		bodyLeft:  container.NewVBox(),
		bodyRight: container.NewVBox(),
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)

	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})

	e.bodyLeft.Add(e.abilities)
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

func (e *FF1Abilities) addAbility() {
	a := &save.Ability{}
	e.c.AddAbility(a)
	e.populate()
}

func (e *FF1Abilities) removeAbility(index int) {
	e.c.RemoveAbility(index)
	e.populate()
}

func (e *FF1Abilities) populate() {
	e.abilities.RemoveAll()
	for i, a := range e.c.Abilities {
		func(i int, a *save.Ability) {
			r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
			e.abilities.Add(container.NewGridWithColumns(4, r.Label, r.ID, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				e.removeAbility(i)
			})))
		}(i, a)
	}
	e.abilities.Add(container.NewGridWithColumns(3, container.NewStack(), container.NewStack(), e.add))
}

func (e *FF1Abilities) addAbilityTable(body fyne.CanvasObject, includeButton bool) fyne.CanvasObject {
	top := container.NewGridWithColumns(4,
		container.NewStack(),
		widget.NewLabel("Ability ID"))
	if includeButton {
		top.Add(e.add)
	}
	return container.NewBorder(top, nil, nil, nil, body)
}

func (e *FF1Abilities) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Abilities
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(4,
			container.NewVScroll(e.bodyLeft),
			container.NewVScroll(e.bodyRight),
			search.Fields(),
			search.Filter())))
}
