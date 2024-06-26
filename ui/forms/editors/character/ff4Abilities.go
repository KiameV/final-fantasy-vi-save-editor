package character

import (
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
	FF4Abilities struct {
		widget.BaseWidget
		c    *core.Character
		rows []*ff4AbRow
		add  *widget.Button
	}
	ff4AbRow struct {
		*inputs.IdEntry
		remove *widget.Button
	}
)

func NewFF4Abilities(c *core.Character) *FF4Abilities {
	e := &FF4Abilities{
		c: c,
	}
	e.ExtendBaseWidget(e)
	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})
	e.populate()
	return e
}

func (e *FF4Abilities) addAbility() {
	a := &save.Ability{}
	e.c.AddAbility(a)
	e.populate()
}

func (e *FF4Abilities) removeAbility(index int) {
	e.c.RemoveAbility(index)
	e.populate()
}

func (e *FF4Abilities) populate() {
	e.rows = make([]*ff4AbRow, len(e.c.Abilities))
	for i, a := range e.c.Abilities {
		func(i int, a *save.Ability) {
			e.rows[i] = &ff4AbRow{
				IdEntry: inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities),
				remove: widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
					e.removeAbility(i)
				}),
			}
		}(i, a)
	}
}

func (e *FF4Abilities) CreateRenderer() fyne.WidgetRenderer {
	body := container.NewVBox()
	for _, r := range e.rows {
		body.Add(container.NewGridWithColumns(3, r.Label, r.ID, r.remove))
	}
	search := inputs.GetSearches().Abilities
	left := container.NewBorder(
		container.NewGridWithColumns(3, container.NewStack(), widget.NewLabel("Ability ID"), e.add),
		nil, nil, nil,
		container.NewVScroll(body))
	right := container.NewGridWithColumns(2, search.Fields(), search.Filter())
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, left, right))
}
