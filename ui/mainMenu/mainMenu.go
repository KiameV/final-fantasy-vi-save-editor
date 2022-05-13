package mainMenu

import (
	"ffvi_editor/global"
	"ffvi_editor/ui"
	"ffvi_editor/ui/character"
	"ffvi_editor/ui/cheats"
	"ffvi_editor/ui/espers"
	"ffvi_editor/ui/inventory"
	"ffvi_editor/ui/misc"
	"ffvi_editor/ui/party"
	"ffvi_editor/ui/skills"
	"ffvi_editor/ui/veldt"
	"github.com/aarzilli/nucular"
)

type mainMenu struct {
	uis [8]ui.UI
}

func NewUI() ui.UI {
	return &mainMenu{
		uis: [8]ui.UI{
			character.NewUI(),
			inventory.NewUI(),
			skills.NewUI(),
			espers.NewUI(),
			misc.NewUI(),
			party.NewUI(),
			veldt.NewUI(),
			cheats.NewUI(),
		},
	}
}

func (u *mainMenu) Draw(w *nucular.Window) {
	w.Row(5).Static(1)

	for i := 0; i < len(u.uis); i++ {
		b := u.uis[i].Behavior()
		if b == ui.Hide ||
			(b == ui.PrShow && !global.IsShowingPR()) ||
			(b == ui.SnesShow && global.IsShowingPR()) {
			continue
		}

		if w.TreePush(nucular.TreeTab, u.uis[i].Name(), false) {
			u.uis[i].Draw(w)
			w.TreePop()
		}
	}
}

func (u *mainMenu) Refresh() {
	for _, ui := range u.uis {
		ui.Refresh()
	}
}

func (u *mainMenu) Name() string {
	return "Main Menu"
}

func (u *mainMenu) Behavior() ui.Behavior {
	return ui.Show
}
