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
			inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&data.Player.Position.X), 3),
			inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&data.Player.Position.Y), 3),
			inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&data.Player.Position.Z), 3),
			inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&data.Player.Direction), 3),
		},
		gps: []fyne.CanvasObject{
			inputs.NewLabeledEntry("World", inputs.NewIntEntryWithData(&data.Gps.MapID), 3),
			inputs.NewLabeledEntry("Area ID", inputs.NewIntEntryWithData(&data.Gps.AreaID), 3),
			inputs.NewLabeledEntry("ID", inputs.NewIntEntryWithData(&data.Gps.GpsID), 3),
			inputs.NewLabeledEntry("Width", inputs.NewIntEntryWithData(&data.Gps.Width), 3),
			inputs.NewLabeledEntry("Height", inputs.NewIntEntryWithData(&data.Gps.Height), 3),
		},
		misc: []fyne.CanvasObject{
			inputs.NewLabeledEntry("Map ID", inputs.NewIntEntryWithData(&data.Map.MapID), 3),
			inputs.NewLabeledEntry("Point In", inputs.NewIntEntryWithData(&data.Map.PointIn), 3),
			inputs.NewLabeledEntry("Transportation ID", inputs.NewIntEntryWithData(&data.Map.TransportationID), 3),
			inputs.NewLabeledEntry("Carrying Hover Ship", widget.NewCheckWithData("", binding.BindBool(&data.Map.CarryingHoverShip)), 3),
		},
	}
	e.ExtendBaseWidget(e)
	return e
}

func (e *MapData) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Maps
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil,
		container.NewVScroll(container.NewVBox(
			widget.NewCard("Player", "", container.NewVBox(e.player...)),
			widget.NewCard("GPS", "", container.NewVBox(e.gps...)),
			widget.NewCard("Misc", "", container.NewVBox(e.misc...)))),
		container.NewGridWithColumns(2, search.Fields(), search.Filter())))
}
