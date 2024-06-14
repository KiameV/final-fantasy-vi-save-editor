package editors

import (
	"ffvi_editor/models/pr"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	Party struct {
		widget.BaseWidget
		enabled binding.Bool
		members [4]*widget.Select
	}
)

func NewParty() *Party {
	p := pr.GetParty()
	e := &Party{
		enabled: binding.BindBool(&p.Enabled),
	}
	e.ExtendBaseWidget(e)
	for i, m := range p.Members {
		func(i int, m *pr.Member) {
			e.members[i] = widget.NewSelect(p.PossibleNames, func(s string) {
				p.Members[i] = p.Possible[s]
			})
			e.members[i].SetSelected(m.Name)
		}(i, m)
	}
	return e
}

func (e *Party) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewGridWithColumns(3,
		container.NewVBox(
			widget.NewLabel("Warning: can cause soft locks and crashing in game."),
			widget.NewCheckWithData("Enabled", e.enabled),
			e.members[0],
			e.members[1],
			e.members[2],
			e.members[3],
		)))
}
