package editors

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Party struct {
		widget.BaseWidget
		left  *fyne.Container
		right *fyne.Container
	}
)

func NewCoreParty(p *core.Party, parties *core.Parties) *Party {
	e := &Party{
		BaseWidget: widget.BaseWidget{},
		left:       container.NewVBox(),
		right:      container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	e.left.Add(widget.NewLabel("Current Party"))
	for _, m := range p.Members {
		func(m *save.CorpsSlotInfo) {
			j := inputs.NewIdEntryWithDataWithHint(&m.CharacterID, finder.Characters)
			e.left.Add(container.NewPadded(
				container.NewGridWithColumns(3, j.Label, j.ID)))
		}(m)
	}
	for i, j := range parties.Parties {
		if i > 0 {
			e.right.Add(widget.NewSeparator())
		}
		e.right.Add(widget.NewLabel(fmt.Sprintf("Party %d", i+1)))
		for _, m := range j {
			in := inputs.NewIdEntryWithDataWithHint(&m.CharacterID, finder.Characters)
			e.right.Add(container.NewPadded(
				container.NewGridWithColumns(3, in.Label, in.ID)))
		}
	}
	return e
}

func (e *Party) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Characters
	return widget.NewSimpleRenderer(container.NewGridWithColumns(2,
		container.NewVScroll(container.NewGridWithColumns(3,
			e.left,
			container.NewVScroll(e.right),
			search.Fields()))))
}
