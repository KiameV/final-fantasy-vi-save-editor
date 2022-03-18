package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
	pm "ffvi_editor/models/pr"
	sm "ffvi_editor/models/snes"
	"ffvi_editor/ui"
	"ffvi_editor/ui/file"
	"fmt"
	"github.com/aarzilli/nucular"
)

var character *models.Character

type widget interface {
	Draw(w *nucular.Window)
	Update(character *models.Character)
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

func NewUI() ui.UI {
	u := &characterUI{
		characterIndex: 0,
		expandAll:      false,
		stats:          newStatsUI(),
		magic:          newMagicUI(),
		equipment:      newEquipmentUI(),
		commands:       newCommandUI(),
		statusEffects:  newStatusEffectsUI(),
	}
	u.Refresh()
	return u
}

func (u *characterUI) Draw(w *nucular.Window) {
	w.Row(18).Static(100, 10, 100)
	if global.IsShowingPR() {
		if i := w.ComboSimple(pr.Characters, u.characterIndex, 12); i != u.characterIndex {
			u.characterIndex = i
			u.refreshWithCharacter(pm.GetCharacter(pr.Characters[u.characterIndex]))
		}
	} else {
		if i := w.ComboSimple(snes.Characters, u.characterIndex, 12); i != u.characterIndex {
			u.characterIndex = i
			u.refreshWithCharacter(sm.GetCharacter(snes.Characters[u.characterIndex]))
		}
	}
	w.Spacing(1)
	w.CheckboxText("Auto-Expand All", &u.expandAll)

	if w.TreePush(nucular.TreeTab, u.makeLabel("Stats"), true) {
		u.stats.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, u.makeLabel("Magic"), u.expandAll) {
		u.magic.Draw(w)
		w.TreePop()
	}
	if w.TreePush(nucular.TreeTab, u.makeLabel("Equipment"), u.expandAll) {
		u.equipment.Draw(w)
		w.TreePop()
	}
	if !global.IsShowingPR() {
		if w.TreePush(nucular.TreeTab, u.makeLabel("Commands"), u.expandAll) {
			u.commands.Draw(w)
			w.TreePop()
		}
	}
	if !global.IsShowingPR() {
		if w.TreePush(nucular.TreeTab, u.makeLabel("Status Effects"), u.expandAll) {
			u.statusEffects.Draw(w)
			w.TreePop()
		}
	}
}

func (u *characterUI) Refresh() {
	u.characterIndex = 0
	if global.IsShowingPR() {
		character = pm.GetCharacter(pr.Characters[0])
	} else {
		character = sm.GetCharacter(snes.Characters[0])
	}
	u.refreshWithCharacter(character)
}

func (u *characterUI) refreshWithCharacter(c *models.Character) {
	character = &models.Character{}
	u.stats.Update(c)
	u.magic.Update(c)
	u.equipment.Update(c)
	u.commands.Update(c)
	u.statusEffects.Update(c)
	character = c
}

func (u *characterUI) Name() string {
	return "Characters"
}

func (u *characterUI) IsPRSupported() bool {
	return true
}

func (u *characterUI) makeLabel(label string) string {
	name := character.Name
	if global.IsShowingPR() && file.PrIO != nil && file.PrIO.HasUnicodeNames() {
		name = pr.Characters[u.characterIndex]
	}
	return fmt.Sprintf("%s - %s", label, name)
}
