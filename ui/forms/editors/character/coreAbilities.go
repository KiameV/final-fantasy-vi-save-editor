package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Abilities struct {
		widget.BaseWidget
		c         *core.Character
		abilities *fyne.Container
		add       *widget.Button
	}
)

func NewCoreAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		c:         c,
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})
	e.populate()
	return e
}

func (e *Abilities) addAbility() {
	a := &core.Ability{}
	e.c.AddAbility(a)
	e.populate()
}

func (e *Abilities) removeAbility(index int) {
	e.c.RemoveAbility(index)
	e.populate()
}

func (e *Abilities) populate() {
	e.abilities.RemoveAll()
	for i, a := range e.c.Abilities {
		func(i int, a *core.Ability) {
			r := inputs.NewIdCountEntryWithDataWithHint(&a.ID, &a.SkillLevel, finder.Abilities)
			in := inputs.NewIntEntryWithData(&a.ContentID)
			e.abilities.Add(container.NewGridWithColumns(5, r.Label, r.ID, in, r.Count, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				e.removeAbility(i)
			})))
		}(i, a)
	}
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(2,
			container.NewBorder(
				container.NewGridWithColumns(5,
					container.NewStack(),
					widget.NewLabel("Ability ID"),
					widget.NewLabel("Content ID"),
					widget.NewLabel("Level"),
				), nil, nil, nil,
				container.NewVScroll(e.abilities)))))
}
