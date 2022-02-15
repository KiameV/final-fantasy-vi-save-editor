package character

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

var character *models.Character

type widget interface {
	Draw(w *nucular.Window)
	Update()
}

type characterUI struct {
	characterIndex int
	expandAll      bool
	stats          widget
	magic          widget
	equipment      widget
	commands       widget
	statusEffects  widget
}

func NewCharacterUI() ui.UI {
	character = models.GetCharacter(consts.Characters[0])
	return &characterUI{
		characterIndex: 0,
		expandAll:      false,
		stats:          newStatsUI(),
		magic:          newMagicUI(),
		equipment:      newEquipmentUI(),
		commands:       newCommandUI(),
		statusEffects:  newStatusEffectsUI(),
	}
}

func (u *characterUI) Draw(w *nucular.Window) {
	w.Row(18).Static(100, 10, 100)
	if i := w.ComboSimple(consts.Characters, u.characterIndex, 12); i != u.characterIndex {
		u.characterIndex = i
		character = models.GetCharacter(consts.Characters[u.characterIndex])
		u.stats.Update()
	}
	w.Spacing(1)
	w.CheckboxText("Auto-Expand All", &u.expandAll)

	if w.TreePush(nucular.TreeTab, "Stats - "+character.Name, true) {
		u.stats.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, "Magic - "+character.Name, u.expandAll) {
		u.magic.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, "Equipment - "+character.Name, u.expandAll) {
		u.equipment.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, "Commands - "+character.Name, u.expandAll) {
		u.commands.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, "Status Effects - "+character.Name, u.expandAll) {
		u.statusEffects.Draw(w)
		w.TreePop()
	}
}