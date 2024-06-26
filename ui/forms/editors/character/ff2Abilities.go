package character

import (
	"strconv"
	"strings"

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
	FF2Abilities struct {
		widget.BaseWidget
		c         *core.Character
		abilities *fyne.Container
		add       *widget.Button
	}
)

func NewFF2Abilities(c *core.Character) *FF2Abilities {
	e := &FF2Abilities{
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

func (e *FF2Abilities) addAbility() {
	e.c.AddAbility(&save.Ability{})
	e.populate()
}

func (e *FF2Abilities) removeAbility(index int) {
	e.c.RemoveAbility(index)
	e.populate()
}

func (e *FF2Abilities) populate() {
	e.abilities.RemoveAll()
	for i, a := range e.c.Abilities {
		func(i int, a *save.Ability) {
			lvl := inputs.NewIntEntryWithData(&a.SkillLevel)
			id := inputs.NewIdEntryWithDataWithHintWithChange(&a.AbilityID, finder.Abilities, func(s string) {
				if j, err := strconv.Atoi(s); err == nil && j > 0 {
					if namelvl, ok := finder.Abilities(j); ok {
						l, _ := strconv.Atoi(strings.Split(namelvl, " ")[1])
						var k int
						if l <= 5 {
							k = l * 25
						} else if l <= 10 {
							k = (l-5)*50 + 100
						} else {
							k = (l-10)*75 + 300
						}
						if k > l {
							lvl.Entry.SetText(strconv.Itoa(k))
						}
					}
				}
			})
			e.abilities.Add(container.NewGridWithColumns(4,
				id.Label, id.ID, lvl, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
					e.removeAbility(i)
				})))
		}(i, a)
	}
}

func (e *FF2Abilities) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Abilities
	left := container.NewBorder(
		container.NewGridWithColumns(4, container.NewStack(), widget.NewLabel("Ability ID"), widget.NewLabel("Skill Level"), e.add),
		nil, nil, nil,
		container.NewVScroll(e.abilities))
	right := container.NewGridWithColumns(2, search.Fields(), search.Filter())
	return widget.NewSimpleRenderer(
		container.NewBorder(
			nil, nil, left, right, nil,
		))
}
