package transportation

import (
	"ffvi_editor/global"
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
	"math"
)

type transportationUI struct{}

func NewUI() ui.UI {
	return &transportationUI{}
}

func (u *transportationUI) Draw(w *nucular.Window) {
	w.Row(24).Static(200)
	w.Label("Can cause soft locks!", "LC")
	if len(pr.Transportations) >= 5 {
		u.draw(w, pr.Transportations[3], "Blackjack")
		w.Row(4).Static()
		u.draw(w, pr.Transportations[4], "Falcon")
	}
}

func (u *transportationUI) draw(w *nucular.Window, t *pr.Transportation, name string) {
	w.Row(24).Static(100, 10, 100)
	w.Label(name, "LC")
	w.Spacing(1)
	if w.CheckboxText("Enabled", &t.Enabled) {
		if t.Enabled {
			t.ForcedEnabled = true
			t.ForcedDisabled = false
		} else {
			t.ForcedEnabled = false
			t.ForcedDisabled = true
		}
	}

	if t.Enabled {
		w.Row(24).Static(200)
		w.PropertyInt("Map ID", math.MinInt, &t.MapID, math.MaxInt, 1, 0)

		w.Row(24).Static(200)
		w.PropertyFloat("Position X", -math.MaxFloat32, &t.Position.X, math.MaxFloat32, 0.1, 0, 5)
		w.Row(24).Static(200)
		w.PropertyFloat("Position Y", -math.MaxFloat32, &t.Position.Y, math.MaxFloat32, 0.1, 0, 5)
		w.Row(24).Static(200)
		w.PropertyFloat("Position Z", -math.MaxFloat32, &t.Position.Z, math.MaxFloat32, 0.1, 0, 5)

		w.Row(24).Static(200)
		w.PropertyInt("Direction", 0, &t.Direction, math.MaxInt, 1, 0)

		w.Row(24).Static(300, 10, 50)
		f := float64(t.TimeStampTicks)
		if w.PropertyFloat("Timestamp Ticks", 0, &f, math.MaxUint64, 1, 0, 100) {
			t.TimeStampTicks = uint64(f)
		}
		w.Spacing(1)
		if w.ButtonText("now") {
			t.TimeStampTicks = global.NowToTicks()
		}
	}
}

func (u *transportationUI) Refresh() {

}

func (u *transportationUI) Name() string {
	return "Transportation"
}

func (u *transportationUI) Behavior() ui.Behavior {
	return ui.PrShow
}
