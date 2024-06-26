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
			c, _ = save.Characters.GetByName(name)
			tabs = container.NewAppTabs(
				container.NewTabItem("Stats", character.NewCoreStats(c)),
				container.NewTabItem("Abilities", s.abilities(game, c)),
				container.NewTabItem("Equipment", character.NewCoreEquipment(c)),
				container.NewTabItem("Commands", character.NewCoreCommands(c)),
			)
		)
		if game == global.Five {
			tabs.Append(container.NewTabItem("Jobs", NewFF5Jobs(c)))
		}
		s.middle.RemoveAll()
		s.middle.Add(tabs)
	}), 2))
	return s
}

func (s *Characters) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewBorder(s.top, nil, s.middle, nil))
}

func (s *Characters) abilities(game global.Game, c *core.Character) (a fyne.CanvasObject) {
	if game == global.One || game == global.Three {
		a = character.NewFF13Abilities(c)
	} else if game == global.Two {
		a = character.NewFF2Abilities(c)
	} else if game == global.Four {
		a = character.NewFF4Abilities(c)
	} else {
		a = container.NewStack()
	}
	return
}
