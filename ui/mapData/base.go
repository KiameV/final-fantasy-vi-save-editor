package mapData

import (
	"ffvi_editor/models/consts"
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
	"math"
	"strings"
)

type mapDataUI struct {
	searchTB nucular.TextEditor
	mapIDsTB nucular.TextEditor
}

func NewUI() ui.UI {
	u := &mapDataUI{}
	u.searchTB.Flags = nucular.EditField
	u.searchTB.Maxlen = 10
	u.searchTB.SingleLine = true

	u.mapIDsTB.Flags = nucular.EditBox | nucular.EditMultiline
	u.mapIDsTB.SingleLine = false
	u.mapIDsTB.Text([]rune(consts.MapText))
	u.mapIDsTB.Flags |= nucular.EditReadOnly
	u.mapIDsTB.Flags |= nucular.EditNoHorizontalScroll
	return u
}

func (u *mapDataUI) Draw(w *nucular.Window) {
	m := pr.GetMapData()

	w.Row(375).Static(250, 400)
	if sw := w.GroupBegin("mapData", 0); sw != nil {
		sw.Row(24).Static(200)
		sw.PropertyInt("Map ID", 1, &m.MapID, len(consts.Maps)-1, 1, 0)
		sw.Row(24).Static(200)
		name := "[invalid]"
		if consts.Maps[m.MapID] != nil {
			name = consts.Maps[m.MapID].String()
		}
		sw.Label(name, "LC")

		//sw.Row(24).Static(410)
		//sw.PropertyInt("Playable Character Corps Id", 0, &m.PlayableCharacterCorpsID, math.MaxInt, 1000, 0)

		sw.Row(24).Static(200)
		sw.Label("Player Position:", "LC")

		sw.Row(24).Static(200)
		sw.PropertyFloat("X", -math.MaxFloat32, &m.Player.X, math.MaxFloat32, 0.1, 0, 5)
		sw.Row(24).Static(200)
		sw.PropertyFloat("Y", -math.MaxFloat32, &m.Player.Y, math.MaxFloat32, 0.1, 0, 5)
		sw.Row(24).Static(200)
		sw.PropertyFloat("Z", -math.MaxFloat32, &m.Player.Z, math.MaxFloat32, 0.1, 0, 5)
		sw.Row(24).Static(200)
		sw.PropertyInt("Player Direction", 0, &m.PlayerDirection, math.MaxInt, 1, 0)

		sw.Row(5).Static()

		sw.Row(24).Static(200)
		sw.Label("GPS Data:", "LC")

		sw.Row(24).Static(200)
		sw.PropertyInt("Map ID", 0, &m.Gps.MapID, math.MaxInt, 1, 0)
		sw.Row(24).Static(200)
		sw.PropertyInt("Area ID", 0, &m.Gps.AreaID, math.MaxInt, 1, 0)
		sw.Row(24).Static(200)
		sw.PropertyInt("GPS ID", 0, &m.Gps.GpsID, math.MaxInt, 1, 0)
		sw.Row(24).Static(200)
		sw.PropertyInt("Width", 0, &m.Gps.Width, math.MaxInt, 1, 0)
		sw.Row(24).Static(200)
		sw.PropertyInt("Height", 0, &m.Gps.Height, math.MaxInt, 1, 0)

		sw.GroupEnd()
	}

	if sw := w.GroupBegin("mapIDs1", 0); sw != nil {
		sw.Row(200).Static(375)
		u.mapIDsTB.Edit(sw)

		sw.Row(24).Static(50, 150)
		sw.Label("Search:", "LC")
		u.searchTB.Edit(sw)

		if len(u.searchTB.Buffer) > 1 {
			s := strings.ToLower(string(u.searchTB.Buffer))
			for _, i := range consts.Maps {
				if i == nil {
					continue
				}
				if strings.Contains(strings.ToLower(i.Name), s) {
					sw.Row(24).Static(200)
					sw.Label(i.String(), "LC")
				}
			}
		}
		sw.GroupEnd()
	}
}

func (u *mapDataUI) Refresh() {

}

func (u *mapDataUI) Name() string {
	return "Map Data"
}

func (u *mapDataUI) Behavior() ui.Behavior {
	return ui.PrShow
}
