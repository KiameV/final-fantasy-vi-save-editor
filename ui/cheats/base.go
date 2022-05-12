package cheats

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"fmt"
	"github.com/aarzilli/nucular"
)

var EncounterLabels []string

type cheatUI struct{}

func NewUI() ui.UI {
	return &cheatUI{}
}

func (u *cheatUI) Draw(w *nucular.Window) {
	c := pr.GetCheats()
	w.Row(24).Static(300)
	w.Label("Currently In Testing", "LC")
	w.CheckboxText("Enable", &c.Enabled)

	w.Row(24).Static(0)

	w.Row(24).Static(250, 400)
	w.PropertyInt("Chests Opened", 0, &c.OpenedChestCount, 214, 1, 0)
	w.Label("Set to 218 for steam achievement", "LC")

	w.Row(24).Static(300)
	w.CheckboxText("Clear Flag", &c.ClearFlag)
	w.CheckboxText("Is Complete Flag:", &c.IsCompleteFlag)

	w.Row(24).Static(0)

	w.Row(24).Static(250, 400)
	w.Label("Veldt Encounters", "LC")
	w.Row(24).Static(100, 100)
	if w.ButtonText("All True") {
		for i := 0; i < len(c.Encounters); i++ {
			c.Encounters[i] = true
		}
	}
	if w.ButtonText("All False") {
		for i := 0; i < len(c.Encounters); i++ {
			c.Encounters[i] = false
		}
	}

	for i := 0; i < len(c.Encounters); i++ {
		if i < len(EncounterLabels) {
			w.Row(24).Static(30, 30, 200)
			w.Label(fmt.Sprintf("%d", i), "LC")
			w.CheckboxText("", &c.Encounters[i])
			w.Label(EncounterLabels[i], "LC")
		} else {
			w.Row(24).Static(30, 30)
			w.Label(fmt.Sprintf("%d", i), "LC")
			w.CheckboxText("", &c.Encounters[i])
		}
	}
}

func (u *cheatUI) Refresh() {}

func (u *cheatUI) Name() string {
	return "Cheats"
}

func (u *cheatUI) Behavior() ui.Behavior {
	return ui.PrShow
}
