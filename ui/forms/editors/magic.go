package editors

import (
	"strings"

	"ffvi_editor/models"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Magic struct {
		widget.BaseWidget
		search              *widget.Entry
		left, middle, right *fyne.Container
		learn, unlearn      *widget.Button
		c                   *models.Character
	}
)

func NewMagic(c *models.Character) *Magic {
	e := &Magic{
		search: widget.NewEntry(),
		left:   container.NewVBox(),
		middle: container.NewVBox(),
		right:  container.NewVBox(),
		c:      c,
	}
	e.ExtendBaseWidget(e)

	bindings := make([]inputs.IntEntryBinding, len(c.SpellsSorted))
	spells := make(map[string]fyne.CanvasObject)
	include := make([]fyne.CanvasObject, len(c.SpellsSorted))
	for i, s := range c.SpellsSorted {
		b := inputs.NewIntEntryBinding(&s.Value)
		in := inputs.NewIntEntryWithBinding(b)
		bindings[i] = b
		spells[s.Name] = inputs.NewLabeledEntry(s.Name, in)
		include[i] = spells[s.Name]
	}
	e.populate(include)

	e.search.OnChanged = func(s string) {
		include = include[:0]
		text := strings.ToLower(e.search.Text)
		for _, v := range c.SpellsSorted {
			if strings.Contains(strings.ToLower(v.Name), text) {
				include = append(include, spells[v.Name])
			}
		}
		e.populate(include)
	}
	e.learn = widget.NewButton("Learn All", func() {
		for _, v := range bindings {
			v.Set(100)
		}
	})
	e.unlearn = widget.NewButton("Unlearn All", func() {
		for _, v := range bindings {
			v.Set(0)
		}
	})
	return e
}

func (e *Magic) populate(include []fyne.CanvasObject) {
	e.left.RemoveAll()
	e.middle.RemoveAll()
	e.right.RemoveAll()
	var s1 int
	if l := len(include); l <= 6 {
		s1 = l + 1
	} else {
		s1 = len(include) / 3
	}
	s2 := 2 * s1
	s1 += 1
	s2 += 1
	for i, v := range include {
		if i < s1 {
			e.left.Add(v)
		} else if i < s2 {
			e.middle.Add(v)
		} else {
			e.right.Add(v)
		}
	}
}

func (e *Magic) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(
		container.NewGridWithColumns(5,
			inputs.NewLabeledEntry("Search:", e.search),
			container.NewPadded(e.learn),
			container.NewPadded(e.unlearn)), nil, nil, nil,
		container.NewVScroll(
			container.NewGridWithColumns(4, e.left, e.middle, e.right))))
}
