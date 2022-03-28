package cheats

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

type cheatUI struct{}

func NewUI() ui.UI {
	return &cheatUI{}
}

func (u *cheatUI) Draw(w *nucular.Window) {
	c := pr.GetCheats()
	w.Row(24).Static(300)
	w.Label("Currently In Testing", "LC")
	w.CheckboxText("Enable", &c.Enabled)

	w.Row(24).Static(250, 400)
	w.CheckboxText("Set All Encounters to True", &c.SetAllEncounters)
	w.Label("May work? Does not update steam.", "LC")

	w.Row(24).Static(250, 400)
	w.PropertyInt("Chests Opened", 0, &c.OpenedChestCount, 214, 1, 0)
	w.Label("Does not update steam. Maybe need to open one more chest?", "LC")

	w.Row(24).Static(300)
	w.CheckboxText("Clear Flag", &c.ClearFlag)
	w.CheckboxText("Is Complete Flag:", &c.IsCompleteFlag)
}

func (u *cheatUI) Refresh() {}

func (u *cheatUI) Name() string {
	return "Cheats"
}

func (u *cheatUI) Behavior() ui.Behavior {
	return ui.PrShow
}
