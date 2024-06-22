package character

import (
	"fmt"

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
		bodyLeft  *fyne.Container
		bodyRight *fyne.Container
		add       *widget.Button
	}
)

func NewCoreAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		c:         c,
		bodyLeft:  container.NewVBox(),
		bodyRight: container.NewVBox(),
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})
	e.bodyLeft.Add(e.addAbilityTable(e.abilities, false))
	e.populate()
	return e
}

func (e *Abilities) FF1() {
	e.bodyLeft.Add(e.addAbilityTable(e.abilities, true))
	for i, lvl := range e.c.FF1.TrainedAbilities {
		l := container.NewGridWithRows(4)
		for _, a := range lvl {
			func(a *core.Ability) {
				r := inputs.NewIdCountEntryWithDataWithHint(&a.ID, &a.SkillLevel, finder.Abilities)
				l.Add(container.NewGridWithColumns(4, r.Label, r.ID, r.Count))
			}(a)
		}
		e.bodyRight.Add(widget.NewCard(fmt.Sprintf("%d:", i+1), "", l))
	}
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
			e.abilities.Add(container.NewGridWithColumns(4, r.Label, r.ID, r.Count, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				e.removeAbility(i)
			})))
		}(i, a)
	}
}

func (e *Abilities) addAbilityTable(body fyne.CanvasObject, includeButton bool) fyne.CanvasObject {
	top := container.NewGridWithColumns(4,
		container.NewStack(),
		widget.NewLabel("Ability ID"),
		widget.NewLabel("Level"))
	if includeButton {
		top.Add(e.add)
	}
	return container.NewBorder(top, nil, nil, nil, container.NewVScroll(body))
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(
		container.NewGridWithColumns(2, e.bodyLeft, e.bodyRight)))
}
