package editors

import (
	"fmt"
	"slices"
	"strings"

	"ffvi_editor/models/pr"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	MapData struct {
		widget.BaseWidget
		search  *widget.Entry
		results *widget.TextGrid
	}
)

func NewMapData() *MapData {
	e := &MapData{
		search:  widget.NewEntry(),
		results: widget.NewTextGrid(),
	}
	e.ExtendBaseWidget(e)

	e.search.OnChanged = func(s string) {
		s = strings.ToLower(s)
		if len(s) < 2 {
			e.results.SetText("")
			return
		}
		var sb strings.Builder
		if len(s) > 2 {
			var found []string
			for k, v := range mapLookup {
				if strings.Contains(strings.ToLower(v), s) {
					found = append(found, fmt.Sprintf("%3d - %s\n", k, v))
				}
			}
			slices.Sort(found)
			for _, v := range found {
				sb.WriteString(v)
			}
		}
		e.results.SetText(sb.String())
	}
	return e
}

func (e *MapData) CreateRenderer() fyne.WidgetRenderer {
	data := pr.GetMapData()
	transport := pr.Transportations

	cards := make([]fyne.CanvasObject, 0, 4)
	cards = append(cards, widget.NewCard("Player", "", container.NewVBox(
		inputs.NewLabeledIntEntryWithHint("Map ID", inputs.NewIntEntryWithData(&data.MapID), inputs.HintArgs{
			Align: inputs.NewAlign(fyne.TextAlignTrailing),
			Hints: &mapLookup,
		}),
		inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&data.Player.X)),
		inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&data.Player.Y)),
		inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&data.Player.Z)),
		inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&data.PlayerDirection)))))
	cards = append(cards,
		widget.NewCard("GPS", "", container.NewVBox(
			inputs.NewLabeledEntry("World", e.newWorldSelect(&data.Gps.MapID)),
			inputs.NewLabeledEntry("Area ID", inputs.NewIntEntryWithData(&data.Gps.AreaID)),
			inputs.NewLabeledEntry("ID", inputs.NewIntEntryWithData(&data.Gps.GpsID)),
			inputs.NewLabeledEntry("Width", inputs.NewIntEntryWithData(&data.Gps.Width)),
			inputs.NewLabeledEntry("Height", inputs.NewIntEntryWithData(&data.Gps.Height)),
		)))
	if len(transport) >= 5 {
		bj := transport[3]
		cards = append(cards,
			widget.NewCard("Blackjack", "", container.NewVBox(
				inputs.NewLabeledEntry("Enabled", widget.NewCheckWithData("", binding.BindBool(&bj.Enabled))),
				inputs.NewLabeledEntry("World", e.newWorldSelectUnset(&bj.MapID)),
				inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&bj.Position.X)),
				inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&bj.Position.Y)),
				inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&bj.Position.Z)),
				inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&bj.Direction)),
			)))
		f := transport[4]
		cards = append(cards,
			widget.NewCard("Falcon", "", container.NewVBox(
				inputs.NewLabeledEntry("Enabled", widget.NewCheckWithData("", binding.BindBool(&f.Enabled))),
				inputs.NewLabeledEntry("World", e.newWorldSelectUnset(&f.MapID)),
				inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&f.Position.X)),
				inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&f.Position.Y)),
				inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&f.Position.Z)),
				inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&f.Direction)),
			)))
		cards = append(cards)
	}

	return widget.NewSimpleRenderer(
		container.NewGridWithColumns(3,
			container.NewVScroll(container.NewVBox(cards...)),
			container.NewBorder(
				widget.NewLabel("Map IDs"), nil, nil, nil,
				container.NewVScroll(mapsTextBox)),
			container.NewBorder(
				inputs.NewLabeledEntry("Find By Name:", e.search), nil, nil, nil,
				container.NewVScroll(e.results))))
}

func (e *MapData) newWorldSelect(i *int) *widget.Select {
	s := widget.NewSelect([]string{"Balance", "Ruin"}, func(s string) {
		if s == "Balance" {
			*i = 1
		} else {
			*i = 2
		}
	})
	if *i == 1 {
		s.SetSelected("Balance")
	} else {
		s.SetSelected("Ruin")
	}
	return s
}

func (e *MapData) newWorldSelectUnset(i *int) *widget.Select {
	s := widget.NewSelect([]string{"-", "Balance", "Ruin"}, func(s string) {
		if s == "Balance" {
			*i = 1
		} else if s == "Ruin" {
			*i = 2
		} else {
			*i = -1
		}
	})
	if *i == 1 {
		s.SetSelected("Balance")
	} else if *i == 2 {
		s.SetSelected("Ruin")
	} else {
		s.SetSelected("-")
	}
	return s
}
