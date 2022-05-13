package cheats

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
	"math"
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

	w.Row(24).Static(0)

	w.Row(24).Static(250, 400)
	_ = w.PropertyInt("Chests Opened", 0, &c.OpenedChestCount, 214, 1, 0)
	w.Label("Set to 218 for steam achievement", "LC")

	w.Row(24).Static(300)
	w.CheckboxText("Clear Flag", &c.ClearFlag)
	w.CheckboxText("Is Complete Flag:", &c.IsCompleteFlag)

	w.Row(24).Static(200)
	w.Label("Played Time", "LC")
	hours, minutes := getTime(int(c.PlayTime))
	w.Row(24).Static(200, 200)
	b1 := w.PropertyInt("Hours", 0, &hours, math.MaxInt, 1, 0)
	b2 := w.PropertyInt("Minutes", 0, &minutes, 59, 1, 0)
	if b1 || b2 {
		c.PlayTime = float64(hours*3600 + minutes*60)
	}
}

func (u *cheatUI) Refresh() {}

func (u *cheatUI) Name() string {
	return "Cheats"
}

func (u *cheatUI) Behavior() ui.Behavior {
	return ui.PrShow
}

func getTime(input int) (hours int, minutes int) {
	hours = int(input / 3600)
	minutes = int(math.Floor(float64(input%(3600)) / 60))
	return
}
