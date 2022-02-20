package misc

import (
	"ffvi_editor/models"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
)

type miscUI struct {
	yLast int
}

func NewUI() ui.UI {
	return &miscUI{}
}

func (u *miscUI) Draw(w *nucular.Window) {
	w.Row(24).Static(200, 10, 200, 10, 250)
	w.PropertyInt("GP:", 0, &models.GetMisc().GP, 16777216, 1000, 0)
	w.Spacing(1)
	w.PropertyInt("Steps:", 0, &models.GetMisc().Steps, 16777216, 1000, 0)
	w.Spacing(1)
	w.PropertyInt("Cursed Shield Fight Count:", 0, &models.GetMisc().CursedShieldFightCount, 255, 1, 0)

	w.Row(5).Static(0)

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Number of Saves:", 0, &models.GetMisc().NumberOfSaves, 255, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Save Count Rollover:", 0, &models.GetMisc().SaveCountRollOver, 255, 1, 0)

	w.Row(5).Static(0)

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Map X Location:", 0, &models.GetMisc().MapXAxis, 255, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Map Y Location:", 0, &models.GetMisc().MapYAxis, 255, 1, 0)

	w.Row(5).Static(0)

	w.Row(24).Static(200)
	w.CheckboxText("Is Airship Visible:", &models.GetMisc().IsAirshipVisible)

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Airship X Loc:", 0, &models.GetMisc().AirshipXAxis, 255, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Airship Y Loc:", 0, &models.GetMisc().AirshipYAxis, 255, 1, 0)
}

func (u *miscUI) Refresh() {

}
