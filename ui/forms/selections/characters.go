package selections

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui/forms/editors"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Characters struct {
		widget.BaseWidget
		top    *fyne.Container
		middle *fyne.Container
	}
)

func NewCharacters() *Characters {
	s := &Characters{
		top:    container.NewHBox(),
		middle: container.NewStack(),
	}
	s.ExtendBaseWidget(s)
	s.top.Add(inputs.NewLabeledEntry("Character:", widget.NewSelect(pr.CharacterNamesHumanSelect(), func(name string) {
		s.middle.RemoveAll()
		c := pr.GetCharacter(name)
		s.middle.Add(container.NewAppTabs(
			container.NewTabItem("Stats", editors.NewCharacter(c)),
			container.NewTabItem("Magic", editors.NewMagic(c)),
			container.NewTabItem("Equipment", editors.NewEquipment(c)),
			container.NewTabItem("Commands", editors.NewCommands(c)),
		))
	})))
	return s
}

func (s *Characters) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(s.top, nil, nil, nil, s.middle))
}
