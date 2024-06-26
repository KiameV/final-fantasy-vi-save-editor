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

func NewCharacters(game global.Game, save *core.Save) *Characters {
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
		if game == global.One || game == global.Three {
			abilities = character.NewFF13Abilities(c)
		} else if game == global.Two {
			abilities = character.NewFF2Abilities(c)
		} else if game == global.Four {
			abilities = character.NewFF4Abilities(c)
		}
		s.middle.RemoveAll()
		s.middle.Add(container.NewAppTabs(
			container.NewTabItem("Stats", character.NewCoreStats(c)),
			container.NewTabItem("Abilities", abilities),
			container.NewTabItem("Equipment", character.NewCoreEquipment(c)),
			container.NewTabItem("Commands", character.NewCoreCommands(c)),
		))
	}), 2))
	return s
}

func (s *Characters) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(s.top, nil, s.middle, nil))
}
