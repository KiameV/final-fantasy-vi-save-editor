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
		current *fyne.Container
		parties *fyne.Container
	}
)

func NewCoreParty(p *core.Party, parties *core.Parties) *Party {
	e := &Party{
		BaseWidget: widget.BaseWidget{},
		current:    container.NewVBox(),
		parties:    container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	e.current.Add(widget.NewLabel("Current Party"))
	for _, m := range p.Members {
		func(m *save.CorpsSlotInfo) {
			j := inputs.NewIdEntryWithDataWithHint(&m.CharacterID, finder.Characters)
			e.current.Add(container.NewPadded(
				container.NewGridWithColumns(3, j.Label, j.ID)))
		}(m)
	}
	for i, j := range parties.Parties {
		if i > 0 {
			e.parties.Add(widget.NewSeparator())
		}
		e.parties.Add(widget.NewLabel(fmt.Sprintf("Party %d", i+1)))
		for _, m := range j {
			in := inputs.NewIdEntryWithDataWithHint(&m.CharacterID, finder.Characters)
			e.parties.Add(container.NewPadded(
				container.NewGridWithColumns(3, in.Label, in.ID)))
		}
	}
	return e
}

func (e *Party) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Characters
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil,
		container.NewVScroll(container.NewGridWithColumns(2,
			e.current,
			container.NewVScroll(e.parties))), search.Fields()))
}
