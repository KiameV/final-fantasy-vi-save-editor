package mapData

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	MapData struct {
		widget.BaseWidget
		player []fyne.CanvasObject
		gps    []fyne.CanvasObject
		misc   []fyne.CanvasObject
	}
)

func NewCore(data *core.MapData) *MapData {
	e := &MapData{
		player: []fyne.CanvasObject{
			inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&data.Player.Position.X)),
			inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&data.Player.Position.Y)),
			inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&data.Player.Position.Z)),
			inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&data.Player.Direction)),
		},
		gps: []fyne.CanvasObject{
			inputs.NewLabeledEntry("World", inputs.NewIntEntryWithData(&data.Gps.MapID)),
			inputs.NewLabeledEntry("Area ID", inputs.NewIntEntryWithData(&data.Gps.AreaID)),
			inputs.NewLabeledEntry("ID", inputs.NewIntEntryWithData(&data.Gps.GpsID)),
			inputs.NewLabeledEntry("Width", inputs.NewIntEntryWithData(&data.Gps.Width)),
			inputs.NewLabeledEntry("Height", inputs.NewIntEntryWithData(&data.Gps.Height)),
		},
		misc: []fyne.CanvasObject{
			inputs.NewLabeledEntry("Map ID", inputs.NewIntEntryWithData(&data.Map.MapID)),
			inputs.NewLabeledEntry("Point In", inputs.NewIntEntryWithData(&data.Map.PointIn)),
			inputs.NewLabeledEntry("Transportation ID", inputs.NewIntEntryWithData(&data.Map.TransportationID)),
			inputs.NewLabeledEntry("Carrying Hover Ship", widget.NewCheckWithData("", binding.BindBool(&data.Map.CarryingHoverShip))),
		},
	}
	e.ExtendBaseWidget(e)
	return e
}

func (e *MapData) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Maps
	return widget.NewSimpleRenderer(container.NewGridWithColumns(2,
		container.NewGridWithColumns(3,
			container.NewVScroll(container.NewVBox(
				widget.NewCard("Player", "", container.NewVBox(e.player...)),
				widget.NewCard("GPS", "", container.NewVBox(e.gps...)),
				widget.NewCard("Misc", "", container.NewVBox(e.misc...)))),
			search.Fields(),
			search.Filter())))
}
