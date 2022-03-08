package mainMenu

import (
	"ffvi_editor/global"
	"ffvi_editor/ui"
	"ffvi_editor/ui/character"
	"ffvi_editor/ui/espers"
	"ffvi_editor/ui/inventory"
	"ffvi_editor/ui/misc"
	"ffvi_editor/ui/skills"
	"github.com/aarzilli/nucular"
)

type mainMenu struct {
	uis [5]ui.UI
}

func NewUI() ui.UI {
	return &mainMenu{
		uis: [5]ui.UI{
			character.NewUI(),
			inventory.NewUI(),
			skills.NewUI(),
			espers.NewUI(),
			misc.NewUI(),
		},
	}
}

func (u *mainMenu) Draw(w *nucular.Window) {
	w.Row(5).Static(1)

	for i := 0; i < len(u.uis); i++ {
		if w.TreePush(nucular.TreeTab, u.uis[i].Name(), false) {
			if global.IsShowingPR() && !u.uis[i].IsPRSupported() {
				w.Row(30).Dynamic(1)
				w.Label("Coming soon for Pixel Remastered", "LC")
			} else {
				u.uis[i].Draw(w)
			}
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

func (u *mainMenu) IsPRSupported() bool {
	return true
}
