package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/editors/character"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Characters struct {
		widget.BaseWidget
		top    *fyne.Container
		middle *fyne.Container
	}
)

func NewCharacters(game global.Game, save *core.Save, allSearch *inputs.Search) *Characters {
	s := &Characters{
		top:    container.NewHBox(),
		middle: container.NewStack(),
	}
	s.ExtendBaseWidget(s)
	s.top.Add(inputs.NewLabeledEntry("Character:", widget.NewSelect(save.Characters.Names(), func(name string) {
		var (
			c, _      = save.Characters.GetByName(name)
			abilities fyne.CanvasObject
		)
		if game == global.One {
			abilities = character.NewFF1Abilities(c)
		} else {
			abilities = character.NewCoreAbilities(c)
		}
		s.middle.RemoveAll()
		s.middle.Add(container.NewAppTabs(
			container.NewTabItem("Stats", character.NewCoreStats(c)),
			container.NewTabItem("Abilities", abilities),
			container.NewTabItem("Equipment", character.NewCoreEquipment(c, allSearch)),
			container.NewTabItem("Commands", character.NewCoreCommands(c)),
		))
	})))
	return s
}

func (s *Characters) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(s.top, nil, nil, nil, s.middle))
}
